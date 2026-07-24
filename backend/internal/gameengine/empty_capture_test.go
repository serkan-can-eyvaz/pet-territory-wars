package gameengine

import (
	"testing"

	"github.com/serkan-can-eyvaz/pet-territory-wars/backend/internal/domain/id"
)

func TestResolveEmptyCaptureQualifiedVisit(t *testing.T) {
	t.Parallel()

	input := ResolveWalkInput{PlayerID: id.PlayerID{1}}
	state := HexState{HexID: id.NewHexID(42), Dominance: 0, Version: 7}
	movementRules := MovementRules{
		MinHexPresenceSeconds: 30,
		MinHexDistanceMeters:  20,
	}
	territoryRules := TerritoryRules{InitialDominance: 15}

	testCases := []struct {
		name  string
		visit VisitedHex
	}{
		{
			name:  "presence at threshold",
			visit: VisitedHex{PresenceSeconds: 30, DistanceMeters: 19},
		},
		{
			name:  "distance at threshold",
			visit: VisitedHex{PresenceSeconds: 29, DistanceMeters: 20},
		},
		{
			name:  "only presence satisfies threshold",
			visit: VisitedHex{PresenceSeconds: 31, DistanceMeters: 19},
		},
		{
			name:  "only distance satisfies threshold",
			visit: VisitedHex{PresenceSeconds: 29, DistanceMeters: 21},
		},
		{
			name:  "both satisfy threshold",
			visit: VisitedHex{PresenceSeconds: 31, DistanceMeters: 21},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			change := resolveEmptyCapture(input, state, testCase.visit, movementRules, territoryRules)

			if change.HexID != state.HexID {
				t.Errorf("HexID = %v, want %v", change.HexID, state.HexID)
			}
			if change.ExpectedVersion != state.Version {
				t.Errorf("ExpectedVersion = %d, want %d", change.ExpectedVersion, state.Version)
			}
			if change.PreviousOwnerID != nil {
				t.Errorf("PreviousOwnerID = %v, want nil", change.PreviousOwnerID)
			}
			if change.NewOwnerID == nil || *change.NewOwnerID != input.PlayerID {
				t.Errorf("NewOwnerID = %v, want %v", change.NewOwnerID, input.PlayerID)
			}
			if change.StoredDominance != state.Dominance {
				t.Errorf("StoredDominance = %d, want %d", change.StoredDominance, state.Dominance)
			}
			if change.EffectiveDominance != state.Dominance {
				t.Errorf("EffectiveDominance = %d, want %d", change.EffectiveDominance, state.Dominance)
			}
			if change.NewDominance != territoryRules.InitialDominance {
				t.Errorf("NewDominance = %d, want %d", change.NewDominance, territoryRules.InitialDominance)
			}
			if change.ChangeType != HexChangeTypeEmptyCapture {
				t.Errorf("ChangeType = %q, want %q", change.ChangeType, HexChangeTypeEmptyCapture)
			}
		})
	}
}

func TestResolveEmptyCaptureNoChangePreservesState(t *testing.T) {
	t.Parallel()

	ownerID := id.PlayerID{2}
	movementRules := MovementRules{
		MinHexPresenceSeconds: 30,
		MinHexDistanceMeters:  20,
	}
	territoryRules := TerritoryRules{InitialDominance: 15}

	testCases := []struct {
		name  string
		state HexState
		visit VisitedHex
	}{
		{
			name:  "unqualified visit",
			state: HexState{HexID: id.NewHexID(42), Dominance: 8, Version: 7},
			visit: VisitedHex{PresenceSeconds: 29, DistanceMeters: 19},
		},
		{
			name:  "owned state",
			state: HexState{HexID: id.NewHexID(43), OwnerID: &ownerID, Dominance: 8, Version: 9},
			visit: VisitedHex{PresenceSeconds: 30, DistanceMeters: 20},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			change := resolveEmptyCapture(
				ResolveWalkInput{PlayerID: id.PlayerID{1}},
				testCase.state,
				testCase.visit,
				movementRules,
				territoryRules,
			)

			if change.HexID != testCase.state.HexID {
				t.Errorf("HexID = %v, want %v", change.HexID, testCase.state.HexID)
			}
			if change.ExpectedVersion != testCase.state.Version {
				t.Errorf("ExpectedVersion = %d, want %d", change.ExpectedVersion, testCase.state.Version)
			}
			if change.PreviousOwnerID != testCase.state.OwnerID {
				t.Errorf("PreviousOwnerID = %v, want %v", change.PreviousOwnerID, testCase.state.OwnerID)
			}
			if change.NewOwnerID != testCase.state.OwnerID {
				t.Errorf("NewOwnerID = %v, want %v", change.NewOwnerID, testCase.state.OwnerID)
			}
			if change.StoredDominance != testCase.state.Dominance {
				t.Errorf("StoredDominance = %d, want %d", change.StoredDominance, testCase.state.Dominance)
			}
			if change.EffectiveDominance != testCase.state.Dominance {
				t.Errorf("EffectiveDominance = %d, want %d", change.EffectiveDominance, testCase.state.Dominance)
			}
			if change.NewDominance != testCase.state.Dominance {
				t.Errorf("NewDominance = %d, want %d", change.NewDominance, testCase.state.Dominance)
			}
			if change.ChangeType != HexChangeTypeNoChange {
				t.Errorf("ChangeType = %q, want %q", change.ChangeType, HexChangeTypeNoChange)
			}
		})
	}
}

func TestResolveEmptyCaptureIsDeterministicAndDoesNotMutateInputs(t *testing.T) {
	t.Parallel()

	input := ResolveWalkInput{
		PlayerID: id.PlayerID{1},
		Route: []LocationPoint{
			{Sequence: 1, Latitude: 41.0, Longitude: 29.0},
		},
		ExistingHexes: map[id.HexID]HexState{
			id.NewHexID(42): {HexID: id.NewHexID(42), Dominance: 0, Version: 7},
		},
	}
	state := HexState{HexID: id.NewHexID(42), Dominance: 0, Version: 7}
	visit := VisitedHex{HexID: state.HexID, PresenceSeconds: 30, DistanceMeters: 20}
	movementRules := MovementRules{MinHexPresenceSeconds: 30, MinHexDistanceMeters: 20}
	territoryRules := TerritoryRules{InitialDominance: 15}

	originalPlayerID := input.PlayerID
	originalRoutePoint := input.Route[0]
	originalExistingHex := input.ExistingHexes[state.HexID]
	originalState := state
	originalVisit := visit
	originalMovementRules := movementRules
	originalTerritoryRules := territoryRules

	first := resolveEmptyCapture(input, state, visit, movementRules, territoryRules)
	second := resolveEmptyCapture(input, state, visit, movementRules, territoryRules)

	if first.HexID != second.HexID ||
		first.ExpectedVersion != second.ExpectedVersion ||
		first.StoredDominance != second.StoredDominance ||
		first.EffectiveDominance != second.EffectiveDominance ||
		first.NewDominance != second.NewDominance ||
		first.ChangeType != second.ChangeType ||
		first.PreviousOwnerID != second.PreviousOwnerID ||
		first.NewOwnerID == nil || second.NewOwnerID == nil ||
		*first.NewOwnerID != *second.NewOwnerID {
		t.Error("resolveEmptyCapture returned different results for identical inputs")
	}

	if input.PlayerID != originalPlayerID || input.Route[0] != originalRoutePoint || input.ExistingHexes[state.HexID] != originalExistingHex {
		t.Error("ResolveWalkInput was mutated")
	}
	if state != originalState {
		t.Error("HexState was mutated")
	}
	if visit != originalVisit {
		t.Error("VisitedHex was mutated")
	}
	if movementRules != originalMovementRules {
		t.Error("MovementRules was mutated")
	}
	if territoryRules != originalTerritoryRules {
		t.Error("TerritoryRules was mutated")
	}
}
