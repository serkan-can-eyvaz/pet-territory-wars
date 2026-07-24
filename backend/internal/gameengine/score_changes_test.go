package gameengine

import (
	"testing"

	"github.com/serkan-can-eyvaz/pet-territory-wars/backend/internal/domain/id"
)

func TestResolveScoreChanges(t *testing.T) {
	t.Parallel()

	previousOwnerID := id.PlayerID{1}
	input := ResolveWalkInput{
		WalkID:   id.WalkID{3},
		PlayerID: id.PlayerID{2},
	}

	testCases := []struct {
		name   string
		change HexChange
		want   []ScoreChange
	}{
		{
			name:   "no change",
			change: HexChange{ChangeType: HexChangeTypeNoChange},
			want:   []ScoreChange{},
		},
		{
			name:   "enemy attack",
			change: HexChange{ChangeType: HexChangeTypeEnemyAttack},
			want:   []ScoreChange{},
		},
		{
			name:   "empty capture",
			change: HexChange{ChangeType: HexChangeTypeEmptyCapture},
			want: []ScoreChange{
				{
					PlayerID: input.PlayerID,
					Metric:   ScoreMetricActiveHex,
					Delta:    1,
					Reason:   ScoreChangeReasonEmptyCapture,
					WalkID:   input.WalkID,
				},
				{
					PlayerID: input.PlayerID,
					Metric:   ScoreMetricCapture,
					Delta:    1,
					Reason:   ScoreChangeReasonEmptyCapture,
					WalkID:   input.WalkID,
				},
				{
					PlayerID: input.PlayerID,
					Metric:   ScoreMetricLifetimeCapture,
					Delta:    1,
					Reason:   ScoreChangeReasonEmptyCapture,
					WalkID:   input.WalkID,
				},
			},
		},
		{
			name:   "owner defense",
			change: HexChange{ChangeType: HexChangeTypeOwnerDefense},
			want: []ScoreChange{
				{
					PlayerID: input.PlayerID,
					Metric:   ScoreMetricDefense,
					Delta:    1,
					Reason:   ScoreChangeReasonOwnerDefense,
					WalkID:   input.WalkID,
				},
			},
		},
		{
			name: "ownership transfer",
			change: HexChange{
				ChangeType:      HexChangeTypeOwnershipTransfer,
				PreviousOwnerID: &previousOwnerID,
			},
			want: []ScoreChange{
				{
					PlayerID: input.PlayerID,
					Metric:   ScoreMetricActiveHex,
					Delta:    1,
					Reason:   ScoreChangeReasonOwnershipTransfer,
					WalkID:   input.WalkID,
				},
				{
					PlayerID: input.PlayerID,
					Metric:   ScoreMetricSteal,
					Delta:    1,
					Reason:   ScoreChangeReasonOwnershipTransfer,
					WalkID:   input.WalkID,
				},
				{
					PlayerID: input.PlayerID,
					Metric:   ScoreMetricLifetimeSteal,
					Delta:    1,
					Reason:   ScoreChangeReasonOwnershipTransfer,
					WalkID:   input.WalkID,
				},
				{
					PlayerID: previousOwnerID,
					Metric:   ScoreMetricActiveHex,
					Delta:    -1,
					Reason:   ScoreChangeReasonOwnershipTransfer,
					WalkID:   input.WalkID,
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := resolveScoreChanges(input, testCase.change)

			if actual == nil {
				t.Fatal("resolveScoreChanges returned nil")
			}
			if len(actual) != len(testCase.want) {
				t.Fatalf("len(result) = %d, want %d", len(actual), len(testCase.want))
			}
			for index := range testCase.want {
				if actual[index] != testCase.want[index] {
					t.Errorf("result[%d] = %+v, want %+v", index, actual[index], testCase.want[index])
				}
			}
		})
	}
}

func TestResolveScoreChangesIsDeterministicAndDoesNotMutateInputs(t *testing.T) {
	t.Parallel()

	previousOwnerID := id.PlayerID{1}
	input := ResolveWalkInput{
		WalkID:   id.WalkID{3},
		PlayerID: id.PlayerID{2},
		Route: []LocationPoint{
			{Sequence: 1, Latitude: 41.0, Longitude: 29.0},
		},
		ExistingHexes: map[id.HexID]HexState{
			id.NewHexID(42): {HexID: id.NewHexID(42), OwnerID: &previousOwnerID, Dominance: 40, Version: 7},
		},
	}
	change := HexChange{
		HexID:           id.NewHexID(42),
		PreviousOwnerID: &previousOwnerID,
		ChangeType:      HexChangeTypeOwnershipTransfer,
	}

	originalWalkID := input.WalkID
	originalPlayerID := input.PlayerID
	originalRoutePoint := input.Route[0]
	originalExistingHex := input.ExistingHexes[id.NewHexID(42)]
	originalChange := change

	first := resolveScoreChanges(input, change)
	second := resolveScoreChanges(input, change)

	if len(first) != len(second) {
		t.Fatalf("len(first) = %d, want %d", len(first), len(second))
	}
	for index := range first {
		if first[index] != second[index] {
			t.Errorf("result[%d] differs for identical inputs", index)
		}
	}
	if input.WalkID != originalWalkID || input.PlayerID != originalPlayerID || input.Route[0] != originalRoutePoint || input.ExistingHexes[id.NewHexID(42)] != originalExistingHex {
		t.Error("ResolveWalkInput was mutated")
	}
	if change != originalChange {
		t.Error("HexChange was mutated")
	}
}
