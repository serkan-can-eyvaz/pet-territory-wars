package gameengine

import (
	"testing"
	"time"

	"github.com/serkan-can-eyvaz/pet-territory-wars/backend/internal/domain/id"
)

func TestResolveDomainEvents(t *testing.T) {
	t.Parallel()

	evaluatedAt := time.Date(2026, time.July, 25, 10, 0, 0, 0, time.UTC)
	input := ResolveWalkInput{
		WalkID:        id.WalkID{1},
		PlayerID:      id.PlayerID{2},
		EvaluatedAt:   evaluatedAt,
		EngineVersion: id.EngineVersion("engine-v1"),
	}
	ruleSetVersion := id.RuleSetVersion("rules-v1")

	testCases := []struct {
		name    string
		context domainEventContext
		want    []DomainEvent
	}{
		{
			name: "no change",
			context: domainEventContext{
				Input:          input,
				RuleSetVersion: ruleSetVersion,
				FinalChange:    HexChange{ChangeType: HexChangeTypeNoChange},
			},
			want: []DomainEvent{},
		},
		{
			name: "empty capture with score event",
			context: domainEventContext{
				Input:          input,
				RuleSetVersion: ruleSetVersion,
				FinalChange: HexChange{
					HexID:      id.NewHexID(42),
					ChangeType: HexChangeTypeEmptyCapture,
				},
				ScoreChanges: []ScoreChange{{}},
				EventIDs:     []id.EventID{{3}, {4}},
			},
			want: []DomainEvent{
				{
					ID:             id.EventID{3},
					Type:           EventTypeEmptyHexCaptured,
					OccurredAt:     evaluatedAt,
					WalkID:         input.WalkID,
					PlayerID:       input.PlayerID,
					HexID:          hexIDPointer(id.NewHexID(42)),
					RuleSetVersion: ruleSetVersion,
					EngineVersion:  input.EngineVersion,
				},
				{
					ID:             id.EventID{4},
					Type:           EventTypePlayerScoreChanged,
					OccurredAt:     evaluatedAt,
					WalkID:         input.WalkID,
					PlayerID:       input.PlayerID,
					RuleSetVersion: ruleSetVersion,
					EngineVersion:  input.EngineVersion,
				},
			},
		},
		{
			name: "owner defense",
			context: domainEventContext{
				Input:          input,
				RuleSetVersion: ruleSetVersion,
				FinalChange: HexChange{
					HexID:      id.NewHexID(42),
					ChangeType: HexChangeTypeOwnerDefense,
				},
				EventIDs: []id.EventID{{3}},
			},
			want: []DomainEvent{{
				ID:             id.EventID{3},
				Type:           EventTypeHexDefended,
				OccurredAt:     evaluatedAt,
				WalkID:         input.WalkID,
				PlayerID:       input.PlayerID,
				HexID:          hexIDPointer(id.NewHexID(42)),
				RuleSetVersion: ruleSetVersion,
				EngineVersion:  input.EngineVersion,
			}},
		},
		{
			name: "enemy attack at threat threshold",
			context: domainEventContext{
				Input:          input,
				RuleSetVersion: ruleSetVersion,
				TerritoryRules: TerritoryRules{ThreatThreshold: 10},
				AttackChange:   HexChange{HexID: id.NewHexID(42), NewDominance: 10, ChangeType: HexChangeTypeEnemyAttack},
				FinalChange:    HexChange{ChangeType: HexChangeTypeEnemyAttack},
				EventIDs:       []id.EventID{{3}, {4}},
			},
			want: []DomainEvent{
				{
					ID:             id.EventID{3},
					Type:           EventTypeHexAttacked,
					OccurredAt:     evaluatedAt,
					WalkID:         input.WalkID,
					PlayerID:       input.PlayerID,
					HexID:          hexIDPointer(id.NewHexID(42)),
					RuleSetVersion: ruleSetVersion,
					EngineVersion:  input.EngineVersion,
				},
				{
					ID:             id.EventID{4},
					Type:           EventTypeHexUnderThreat,
					OccurredAt:     evaluatedAt,
					WalkID:         input.WalkID,
					PlayerID:       input.PlayerID,
					HexID:          hexIDPointer(id.NewHexID(42)),
					RuleSetVersion: ruleSetVersion,
					EngineVersion:  input.EngineVersion,
				},
			},
		},
		{
			name: "ownership transfer preserves attack and final contexts",
			context: domainEventContext{
				Input:          input,
				RuleSetVersion: ruleSetVersion,
				AttackChange: HexChange{
					HexID:      id.NewHexID(42),
					ChangeType: HexChangeTypeEnemyAttack,
				},
				FinalChange: HexChange{
					HexID:      id.NewHexID(43),
					ChangeType: HexChangeTypeOwnershipTransfer,
				},
				ScoreChanges: []ScoreChange{{}},
				EventIDs:     []id.EventID{{3}, {4}, {5}},
			},
			want: []DomainEvent{
				{
					ID:             id.EventID{3},
					Type:           EventTypeHexAttacked,
					OccurredAt:     evaluatedAt,
					WalkID:         input.WalkID,
					PlayerID:       input.PlayerID,
					HexID:          hexIDPointer(id.NewHexID(42)),
					RuleSetVersion: ruleSetVersion,
					EngineVersion:  input.EngineVersion,
				},
				{
					ID:             id.EventID{4},
					Type:           EventTypeHexOwnershipTransferred,
					OccurredAt:     evaluatedAt,
					WalkID:         input.WalkID,
					PlayerID:       input.PlayerID,
					HexID:          hexIDPointer(id.NewHexID(43)),
					RuleSetVersion: ruleSetVersion,
					EngineVersion:  input.EngineVersion,
				},
				{
					ID:             id.EventID{5},
					Type:           EventTypePlayerScoreChanged,
					OccurredAt:     evaluatedAt,
					WalkID:         input.WalkID,
					PlayerID:       input.PlayerID,
					RuleSetVersion: ruleSetVersion,
					EngineVersion:  input.EngineVersion,
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := resolveDomainEvents(testCase.context)

			if actual == nil {
				t.Fatal("resolveDomainEvents returned nil")
			}
			if len(actual) != len(testCase.want) {
				t.Fatalf("len(result) = %d, want %d", len(actual), len(testCase.want))
			}
			for index := range testCase.want {
				if actual[index].ID != testCase.want[index].ID ||
					actual[index].Type != testCase.want[index].Type ||
					actual[index].OccurredAt != testCase.want[index].OccurredAt ||
					actual[index].WalkID != testCase.want[index].WalkID ||
					actual[index].PlayerID != testCase.want[index].PlayerID ||
					actual[index].RuleSetVersion != testCase.want[index].RuleSetVersion ||
					actual[index].EngineVersion != testCase.want[index].EngineVersion {
					t.Errorf("result[%d] = %+v, want %+v", index, actual[index], testCase.want[index])
				}
				if actual[index].HexID == nil || testCase.want[index].HexID == nil {
					if actual[index].HexID != testCase.want[index].HexID {
						t.Errorf("result[%d].HexID = %v, want %v", index, actual[index].HexID, testCase.want[index].HexID)
					}
					continue
				}
				if *actual[index].HexID != *testCase.want[index].HexID {
					t.Errorf("result[%d].HexID = %v, want %v", index, *actual[index].HexID, *testCase.want[index].HexID)
				}
			}
		})
	}
}

func TestResolveDomainEventsIsDeterministicAndDoesNotMutateInputs(t *testing.T) {
	t.Parallel()

	ownerID := id.PlayerID{1}
	context := domainEventContext{
		Input: ResolveWalkInput{
			WalkID:        id.WalkID{1},
			PlayerID:      id.PlayerID{2},
			EvaluatedAt:   time.Date(2026, time.July, 25, 10, 0, 0, 0, time.UTC),
			EngineVersion: id.EngineVersion("engine-v1"),
			Route: []LocationPoint{
				{Sequence: 1, Latitude: 41.0, Longitude: 29.0},
			},
			ExistingHexes: map[id.HexID]HexState{
				id.NewHexID(42): {HexID: id.NewHexID(42), OwnerID: &ownerID, Dominance: 40, Version: 7},
			},
		},
		RuleSetVersion: id.RuleSetVersion("rules-v1"),
		AttackChange: HexChange{
			HexID:      id.NewHexID(42),
			ChangeType: HexChangeTypeEnemyAttack,
		},
		FinalChange: HexChange{
			HexID:      id.NewHexID(42),
			ChangeType: HexChangeTypeOwnershipTransfer,
		},
		ScoreChanges: []ScoreChange{{PlayerID: id.PlayerID{2}, Metric: ScoreMetricActiveHex, Delta: 1}},
		EventIDs:     []id.EventID{{3}, {4}, {5}},
	}

	originalWalkID := context.Input.WalkID
	originalPlayerID := context.Input.PlayerID
	originalRoutePoint := context.Input.Route[0]
	originalExistingHex := context.Input.ExistingHexes[id.NewHexID(42)]
	originalAttackChange := context.AttackChange
	originalFinalChange := context.FinalChange
	originalScoreChange := context.ScoreChanges[0]
	originalEventIDs := append([]id.EventID(nil), context.EventIDs...)

	first := resolveDomainEvents(context)
	second := resolveDomainEvents(context)

	if len(first) != len(second) {
		t.Fatalf("len(first) = %d, want %d", len(first), len(second))
	}
	for index := range first {
		if first[index].ID != second[index].ID || first[index].Type != second[index].Type ||
			first[index].WalkID != second[index].WalkID || first[index].PlayerID != second[index].PlayerID ||
			first[index].RuleSetVersion != second[index].RuleSetVersion || first[index].EngineVersion != second[index].EngineVersion {
			t.Errorf("result[%d] differs for identical inputs", index)
		}
	}
	if context.Input.WalkID != originalWalkID || context.Input.PlayerID != originalPlayerID || context.Input.Route[0] != originalRoutePoint || context.Input.ExistingHexes[id.NewHexID(42)] != originalExistingHex {
		t.Error("ResolveWalkInput was mutated")
	}
	if context.AttackChange != originalAttackChange {
		t.Error("AttackChange was mutated")
	}
	if context.FinalChange != originalFinalChange {
		t.Error("FinalChange was mutated")
	}
	if context.ScoreChanges[0] != originalScoreChange {
		t.Error("ScoreChanges was mutated")
	}
	for index := range context.EventIDs {
		if context.EventIDs[index] != originalEventIDs[index] {
			t.Error("EventIDs was mutated")
		}
	}
}

func hexIDPointer(value id.HexID) *id.HexID {
	return &value
}
