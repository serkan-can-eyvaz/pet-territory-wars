package gameengine

import (
	"testing"

	"github.com/serkan-can-eyvaz/pet-territory-wars/backend/internal/domain/id"
)

func TestResolveOwnershipTransfer(t *testing.T) {
	t.Parallel()

	ownerID := id.PlayerID{1}
	attackerID := id.PlayerID{2}
	input := ResolveWalkInput{PlayerID: attackerID}
	state := HexState{HexID: id.NewHexID(42), OwnerID: &ownerID, Dominance: 50, Version: 7}

	testCases := []struct {
		name         string
		state        HexState
		attackChange HexChange
		wantType     HexChangeType
	}{
		{
			name:  "non-attack change",
			state: state,
			attackChange: HexChange{
				ChangeType:         HexChangeTypeNoChange,
				EffectiveDominance: 0,
				NewDominance:       0,
			},
			wantType: HexChangeTypeNoChange,
		},
		{
			name:  "non-zero attack dominance",
			state: state,
			attackChange: HexChange{
				ChangeType:         HexChangeTypeEnemyAttack,
				EffectiveDominance: 10,
				NewDominance:       10,
			},
			wantType: HexChangeTypeNoChange,
		},
		{
			name:  "unowned state",
			state: HexState{HexID: id.NewHexID(42), Dominance: 0, Version: 7},
			attackChange: HexChange{
				ChangeType:         HexChangeTypeEnemyAttack,
				EffectiveDominance: 0,
				NewDominance:       0,
			},
			wantType: HexChangeTypeNoChange,
		},
		{
			name:  "successful transfer",
			state: state,
			attackChange: HexChange{
				ChangeType:         HexChangeTypeEnemyAttack,
				EffectiveDominance: 0,
				NewDominance:       0,
			},
			wantType: HexChangeTypeOwnershipTransfer,
		},
	}

	territoryRules := TerritoryRules{InitialDominance: 40}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			change := resolveOwnershipTransfer(input, testCase.state, testCase.attackChange, territoryRules)

			if change.HexID != testCase.state.HexID {
				t.Errorf("HexID = %v, want %v", change.HexID, testCase.state.HexID)
			}
			if change.ExpectedVersion != testCase.state.Version {
				t.Errorf("ExpectedVersion = %d, want %d", change.ExpectedVersion, testCase.state.Version)
			}
			if change.ChangeType != testCase.wantType {
				t.Errorf("ChangeType = %q, want %q", change.ChangeType, testCase.wantType)
			}

			if testCase.wantType == HexChangeTypeNoChange {
				if change.PreviousOwnerID != testCase.state.OwnerID || change.NewOwnerID != testCase.state.OwnerID {
					t.Errorf("owner IDs = (%v, %v), want (%v, %v)", change.PreviousOwnerID, change.NewOwnerID, testCase.state.OwnerID, testCase.state.OwnerID)
				}
				if change.StoredDominance != testCase.state.Dominance {
					t.Errorf("StoredDominance = %d, want %d", change.StoredDominance, testCase.state.Dominance)
				}
				if change.EffectiveDominance != testCase.attackChange.EffectiveDominance {
					t.Errorf("EffectiveDominance = %d, want %d", change.EffectiveDominance, testCase.attackChange.EffectiveDominance)
				}
				if change.NewDominance != testCase.attackChange.NewDominance {
					t.Errorf("NewDominance = %d, want %d", change.NewDominance, testCase.attackChange.NewDominance)
				}

				return
			}

			if change.PreviousOwnerID != testCase.state.OwnerID {
				t.Errorf("PreviousOwnerID = %v, want %v", change.PreviousOwnerID, testCase.state.OwnerID)
			}
			if change.NewOwnerID == nil || *change.NewOwnerID != input.PlayerID {
				t.Errorf("NewOwnerID = %v, want %v", change.NewOwnerID, input.PlayerID)
			}
			if change.StoredDominance != territoryRules.InitialDominance {
				t.Errorf("StoredDominance = %d, want %d", change.StoredDominance, territoryRules.InitialDominance)
			}
			if change.EffectiveDominance != territoryRules.InitialDominance {
				t.Errorf("EffectiveDominance = %d, want %d", change.EffectiveDominance, territoryRules.InitialDominance)
			}
			if change.NewDominance != territoryRules.InitialDominance {
				t.Errorf("NewDominance = %d, want %d", change.NewDominance, territoryRules.InitialDominance)
			}
		})
	}
}

func TestResolveOwnershipTransferIsDeterministicAndDoesNotMutateInputs(t *testing.T) {
	t.Parallel()

	ownerID := id.PlayerID{1}
	input := ResolveWalkInput{
		PlayerID: id.PlayerID{2},
		Route: []LocationPoint{
			{Sequence: 1, Latitude: 41.0, Longitude: 29.0},
		},
		ExistingHexes: map[id.HexID]HexState{
			id.NewHexID(42): {HexID: id.NewHexID(42), OwnerID: &ownerID, Dominance: 50, Version: 7},
		},
	}
	state := HexState{HexID: id.NewHexID(42), OwnerID: &ownerID, Dominance: 50, Version: 7}
	attackChange := HexChange{
		HexID:              state.HexID,
		ExpectedVersion:    state.Version,
		PreviousOwnerID:    state.OwnerID,
		NewOwnerID:         state.OwnerID,
		StoredDominance:    state.Dominance,
		EffectiveDominance: 0,
		NewDominance:       0,
		ChangeType:         HexChangeTypeEnemyAttack,
	}
	territoryRules := TerritoryRules{InitialDominance: 40}

	originalPlayerID := input.PlayerID
	originalRoutePoint := input.Route[0]
	originalExistingHex := input.ExistingHexes[state.HexID]
	originalState := state
	originalAttackChange := attackChange
	originalTerritoryRules := territoryRules

	first := resolveOwnershipTransfer(input, state, attackChange, territoryRules)
	second := resolveOwnershipTransfer(input, state, attackChange, territoryRules)

	if first.HexID != second.HexID ||
		first.ExpectedVersion != second.ExpectedVersion ||
		first.PreviousOwnerID != second.PreviousOwnerID ||
		first.NewOwnerID == nil || second.NewOwnerID == nil ||
		*first.NewOwnerID != *second.NewOwnerID ||
		first.StoredDominance != second.StoredDominance ||
		first.EffectiveDominance != second.EffectiveDominance ||
		first.NewDominance != second.NewDominance ||
		first.ChangeType != second.ChangeType {
		t.Error("resolveOwnershipTransfer returned different results for identical inputs")
	}
	if input.PlayerID != originalPlayerID || input.Route[0] != originalRoutePoint || input.ExistingHexes[state.HexID] != originalExistingHex {
		t.Error("ResolveWalkInput was mutated")
	}
	if state != originalState {
		t.Error("HexState was mutated")
	}
	if attackChange != originalAttackChange {
		t.Error("HexChange was mutated")
	}
	if territoryRules != originalTerritoryRules {
		t.Error("TerritoryRules was mutated")
	}
}
