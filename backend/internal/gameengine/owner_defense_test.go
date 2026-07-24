package gameengine

import (
	"testing"

	"github.com/serkan-can-eyvaz/pet-territory-wars/backend/internal/domain/id"
)

func TestResolveOwnerDefenseQualifiedVisit(t *testing.T) {
	t.Parallel()

	ownerID := id.PlayerID{1}
	input := ResolveWalkInput{PlayerID: ownerID}
	state := HexState{HexID: id.NewHexID(42), OwnerID: &ownerID, Dominance: 40, Version: 7}
	movementRules := MovementRules{
		MinHexPresenceSeconds: 30,
		MinHexDistanceMeters:  20,
	}

	testCases := []struct {
		name               string
		visit              VisitedHex
		effectiveDominance int
		territoryRules     TerritoryRules
		wantDominance      int
	}{
		{
			name:               "presence at threshold",
			visit:              VisitedHex{PresenceSeconds: 30, DistanceMeters: 19},
			effectiveDominance: 60,
			territoryRules:     TerritoryRules{OwnerVisitGain: 5, MaxDominance: 100},
			wantDominance:      65,
		},
		{
			name:               "distance at threshold",
			visit:              VisitedHex{PresenceSeconds: 29, DistanceMeters: 20},
			effectiveDominance: 60,
			territoryRules:     TerritoryRules{OwnerVisitGain: 5, MaxDominance: 100},
			wantDominance:      65,
		},
		{
			name:               "only presence satisfies threshold",
			visit:              VisitedHex{PresenceSeconds: 31, DistanceMeters: 19},
			effectiveDominance: 60,
			territoryRules:     TerritoryRules{OwnerVisitGain: 5, MaxDominance: 100},
			wantDominance:      65,
		},
		{
			name:               "only distance satisfies threshold",
			visit:              VisitedHex{PresenceSeconds: 29, DistanceMeters: 21},
			effectiveDominance: 60,
			territoryRules:     TerritoryRules{OwnerVisitGain: 5, MaxDominance: 100},
			wantDominance:      65,
		},
		{
			name:               "maximum dominance clamp",
			visit:              VisitedHex{PresenceSeconds: 31, DistanceMeters: 21},
			effectiveDominance: 95,
			territoryRules:     TerritoryRules{OwnerVisitGain: 10, MaxDominance: 100},
			wantDominance:      100,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			change := resolveOwnerDefense(
				input,
				state,
				testCase.visit,
				movementRules,
				testCase.territoryRules,
				testCase.effectiveDominance,
			)

			if change.HexID != state.HexID {
				t.Errorf("HexID = %v, want %v", change.HexID, state.HexID)
			}
			if change.PreviousOwnerID != state.OwnerID || change.NewOwnerID != state.OwnerID {
				t.Errorf("owner IDs = (%v, %v), want (%v, %v)", change.PreviousOwnerID, change.NewOwnerID, state.OwnerID, state.OwnerID)
			}
			if change.StoredDominance != state.Dominance {
				t.Errorf("StoredDominance = %d, want %d", change.StoredDominance, state.Dominance)
			}
			if change.EffectiveDominance != testCase.wantDominance {
				t.Errorf("EffectiveDominance = %d, want %d", change.EffectiveDominance, testCase.wantDominance)
			}
			if change.NewDominance != testCase.wantDominance {
				t.Errorf("NewDominance = %d, want %d", change.NewDominance, testCase.wantDominance)
			}
			if change.ExpectedVersion != state.Version {
				t.Errorf("ExpectedVersion = %d, want %d", change.ExpectedVersion, state.Version)
			}
			if change.ChangeType != HexChangeTypeOwnerDefense {
				t.Errorf("ChangeType = %q, want %q", change.ChangeType, HexChangeTypeOwnerDefense)
			}
		})
	}
}

func TestResolveOwnerDefenseNoChange(t *testing.T) {
	t.Parallel()

	ownerID := id.PlayerID{1}
	otherPlayerID := id.PlayerID{2}
	movementRules := MovementRules{
		MinHexPresenceSeconds: 30,
		MinHexDistanceMeters:  20,
	}
	qualifiedVisit := VisitedHex{PresenceSeconds: 30, DistanceMeters: 20}

	testCases := []struct {
		name               string
		input              ResolveWalkInput
		state              HexState
		visit              VisitedHex
		territoryRules     TerritoryRules
		effectiveDominance int
	}{
		{
			name:               "unowned state",
			input:              ResolveWalkInput{PlayerID: ownerID},
			state:              HexState{HexID: id.NewHexID(42), Dominance: 40, Version: 7},
			visit:              qualifiedVisit,
			territoryRules:     TerritoryRules{OwnerVisitGain: 5, MaxDominance: 100},
			effectiveDominance: 35,
		},
		{
			name:               "different owner",
			input:              ResolveWalkInput{PlayerID: otherPlayerID},
			state:              HexState{HexID: id.NewHexID(42), OwnerID: &ownerID, Dominance: 40, Version: 7},
			visit:              qualifiedVisit,
			territoryRules:     TerritoryRules{OwnerVisitGain: 5, MaxDominance: 100},
			effectiveDominance: 35,
		},
		{
			name:               "unqualified visit",
			input:              ResolveWalkInput{PlayerID: ownerID},
			state:              HexState{HexID: id.NewHexID(42), OwnerID: &ownerID, Dominance: 40, Version: 7},
			visit:              VisitedHex{PresenceSeconds: 29, DistanceMeters: 19},
			territoryRules:     TerritoryRules{OwnerVisitGain: 5, MaxDominance: 100},
			effectiveDominance: 35,
		},
		{
			name:               "zero owner visit gain",
			input:              ResolveWalkInput{PlayerID: ownerID},
			state:              HexState{HexID: id.NewHexID(42), OwnerID: &ownerID, Dominance: 40, Version: 7},
			visit:              qualifiedVisit,
			territoryRules:     TerritoryRules{OwnerVisitGain: 0, MaxDominance: 100},
			effectiveDominance: 35,
		},
		{
			name:               "effective dominance at maximum",
			input:              ResolveWalkInput{PlayerID: ownerID},
			state:              HexState{HexID: id.NewHexID(42), OwnerID: &ownerID, Dominance: 100, Version: 7},
			visit:              qualifiedVisit,
			territoryRules:     TerritoryRules{OwnerVisitGain: 5, MaxDominance: 100},
			effectiveDominance: 100,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			change := resolveOwnerDefense(
				testCase.input,
				testCase.state,
				testCase.visit,
				movementRules,
				testCase.territoryRules,
				testCase.effectiveDominance,
			)

			if change.HexID != testCase.state.HexID {
				t.Errorf("HexID = %v, want %v", change.HexID, testCase.state.HexID)
			}
			if change.PreviousOwnerID != testCase.state.OwnerID || change.NewOwnerID != testCase.state.OwnerID {
				t.Errorf("owner IDs = (%v, %v), want (%v, %v)", change.PreviousOwnerID, change.NewOwnerID, testCase.state.OwnerID, testCase.state.OwnerID)
			}
			if change.StoredDominance != testCase.state.Dominance {
				t.Errorf("StoredDominance = %d, want %d", change.StoredDominance, testCase.state.Dominance)
			}
			if change.EffectiveDominance != testCase.effectiveDominance {
				t.Errorf("EffectiveDominance = %d, want %d", change.EffectiveDominance, testCase.effectiveDominance)
			}
			if change.NewDominance != testCase.effectiveDominance {
				t.Errorf("NewDominance = %d, want %d", change.NewDominance, testCase.effectiveDominance)
			}
			if change.ExpectedVersion != testCase.state.Version {
				t.Errorf("ExpectedVersion = %d, want %d", change.ExpectedVersion, testCase.state.Version)
			}
			if change.ChangeType != HexChangeTypeNoChange {
				t.Errorf("ChangeType = %q, want %q", change.ChangeType, HexChangeTypeNoChange)
			}
		})
	}
}

func TestResolveOwnerDefenseIsDeterministicAndDoesNotMutateInputs(t *testing.T) {
	t.Parallel()

	ownerID := id.PlayerID{1}
	input := ResolveWalkInput{
		PlayerID: ownerID,
		Route: []LocationPoint{
			{Sequence: 1, Latitude: 41.0, Longitude: 29.0},
		},
		ExistingHexes: map[id.HexID]HexState{
			id.NewHexID(42): {HexID: id.NewHexID(42), OwnerID: &ownerID, Dominance: 40, Version: 7},
		},
	}
	state := HexState{HexID: id.NewHexID(42), OwnerID: &ownerID, Dominance: 40, Version: 7}
	visit := VisitedHex{HexID: state.HexID, PresenceSeconds: 30, DistanceMeters: 20}
	movementRules := MovementRules{MinHexPresenceSeconds: 30, MinHexDistanceMeters: 20}
	territoryRules := TerritoryRules{OwnerVisitGain: 5, MaxDominance: 100}
	effectiveDominance := 60

	originalPlayerID := input.PlayerID
	originalRoutePoint := input.Route[0]
	originalExistingHex := input.ExistingHexes[state.HexID]
	originalState := state
	originalVisit := visit
	originalMovementRules := movementRules
	originalTerritoryRules := territoryRules

	first := resolveOwnerDefense(input, state, visit, movementRules, territoryRules, effectiveDominance)
	second := resolveOwnerDefense(input, state, visit, movementRules, territoryRules, effectiveDominance)

	if first != second {
		t.Error("resolveOwnerDefense returned different results for identical inputs")
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
