package gameengine

import (
	"testing"

	"github.com/serkan-can-eyvaz/pet-territory-wars/backend/internal/domain/id"
)

func TestResolveEnemyAttackQualifiedVisit(t *testing.T) {
	t.Parallel()

	ownerID := id.PlayerID{1}
	input := ResolveWalkInput{PlayerID: id.PlayerID{2}}
	state := HexState{HexID: id.NewHexID(42), OwnerID: &ownerID, Dominance: 70, Version: 7}
	movementRules := MovementRules{
		MinHexPresenceSeconds: 30,
		MinHexDistanceMeters:  20,
	}

	testCases := []struct {
		name               string
		visit              VisitedHex
		effectiveDominance int
		damage             int
		wantDominance      int
	}{
		{
			name:               "presence at threshold",
			visit:              VisitedHex{PresenceSeconds: 30, DistanceMeters: 19},
			effectiveDominance: 60,
			damage:             15,
			wantDominance:      45,
		},
		{
			name:               "distance at threshold",
			visit:              VisitedHex{PresenceSeconds: 29, DistanceMeters: 20},
			effectiveDominance: 60,
			damage:             15,
			wantDominance:      45,
		},
		{
			name:               "attack formula",
			visit:              VisitedHex{PresenceSeconds: 31, DistanceMeters: 21},
			effectiveDominance: 70,
			damage:             20,
			wantDominance:      50,
		},
		{
			name:               "damage greater than effective dominance",
			visit:              VisitedHex{PresenceSeconds: 31, DistanceMeters: 21},
			effectiveDominance: 10,
			damage:             15,
			wantDominance:      0,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			change := resolveEnemyAttack(
				input,
				state,
				testCase.visit,
				movementRules,
				TerritoryRules{EnemyAttackDamage: testCase.damage},
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
			if change.ChangeType != HexChangeTypeEnemyAttack {
				t.Errorf("ChangeType = %q, want %q", change.ChangeType, HexChangeTypeEnemyAttack)
			}
		})
	}
}

func TestResolveEnemyAttackNoChange(t *testing.T) {
	t.Parallel()

	ownerID := id.PlayerID{1}
	currentPlayerID := id.PlayerID{2}
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
		damage             int
		effectiveDominance int
	}{
		{
			name:               "unowned state",
			input:              ResolveWalkInput{PlayerID: currentPlayerID},
			state:              HexState{HexID: id.NewHexID(42), Dominance: 40, Version: 7},
			visit:              qualifiedVisit,
			damage:             15,
			effectiveDominance: 35,
		},
		{
			name:               "same owner",
			input:              ResolveWalkInput{PlayerID: ownerID},
			state:              HexState{HexID: id.NewHexID(42), OwnerID: &ownerID, Dominance: 40, Version: 7},
			visit:              qualifiedVisit,
			damage:             15,
			effectiveDominance: 35,
		},
		{
			name:               "unqualified visit",
			input:              ResolveWalkInput{PlayerID: currentPlayerID},
			state:              HexState{HexID: id.NewHexID(42), OwnerID: &ownerID, Dominance: 40, Version: 7},
			visit:              VisitedHex{PresenceSeconds: 29, DistanceMeters: 19},
			damage:             15,
			effectiveDominance: 35,
		},
		{
			name:               "zero attack damage",
			input:              ResolveWalkInput{PlayerID: currentPlayerID},
			state:              HexState{HexID: id.NewHexID(42), OwnerID: &ownerID, Dominance: 40, Version: 7},
			visit:              qualifiedVisit,
			damage:             0,
			effectiveDominance: 35,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			change := resolveEnemyAttack(
				testCase.input,
				testCase.state,
				testCase.visit,
				movementRules,
				TerritoryRules{EnemyAttackDamage: testCase.damage},
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

func TestResolveEnemyAttackIsDeterministicAndDoesNotMutateInputs(t *testing.T) {
	t.Parallel()

	ownerID := id.PlayerID{1}
	input := ResolveWalkInput{
		PlayerID: id.PlayerID{2},
		Route: []LocationPoint{
			{Sequence: 1, Latitude: 41.0, Longitude: 29.0},
		},
		ExistingHexes: map[id.HexID]HexState{
			id.NewHexID(42): {HexID: id.NewHexID(42), OwnerID: &ownerID, Dominance: 70, Version: 7},
		},
	}
	state := HexState{HexID: id.NewHexID(42), OwnerID: &ownerID, Dominance: 70, Version: 7}
	visit := VisitedHex{HexID: state.HexID, PresenceSeconds: 30, DistanceMeters: 20}
	movementRules := MovementRules{MinHexPresenceSeconds: 30, MinHexDistanceMeters: 20}
	territoryRules := TerritoryRules{EnemyAttackDamage: 15}
	effectiveDominance := 60

	originalPlayerID := input.PlayerID
	originalRoutePoint := input.Route[0]
	originalExistingHex := input.ExistingHexes[state.HexID]
	originalState := state
	originalVisit := visit
	originalMovementRules := movementRules
	originalTerritoryRules := territoryRules

	first := resolveEnemyAttack(input, state, visit, movementRules, territoryRules, effectiveDominance)
	second := resolveEnemyAttack(input, state, visit, movementRules, territoryRules, effectiveDominance)

	if first != second {
		t.Error("resolveEnemyAttack returned different results for identical inputs")
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
