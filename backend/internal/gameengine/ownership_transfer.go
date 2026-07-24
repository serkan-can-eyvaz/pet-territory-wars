package gameengine

// resolveOwnershipTransfer calculates the territory change after an enemy attack reaches zero dominance.
func resolveOwnershipTransfer(
	input ResolveWalkInput,
	state HexState,
	attackChange HexChange,
	territoryRules TerritoryRules,
) HexChange {
	change := HexChange{
		HexID:              state.HexID,
		ExpectedVersion:    state.Version,
		PreviousOwnerID:    state.OwnerID,
		NewOwnerID:         state.OwnerID,
		StoredDominance:    state.Dominance,
		EffectiveDominance: attackChange.EffectiveDominance,
		NewDominance:       attackChange.NewDominance,
		ChangeType:         HexChangeTypeNoChange,
	}

	if attackChange.ChangeType != HexChangeTypeEnemyAttack || attackChange.NewDominance != 0 || state.OwnerID == nil {
		return change
	}

	newOwnerID := input.PlayerID
	change.NewOwnerID = &newOwnerID
	change.StoredDominance = territoryRules.InitialDominance
	change.EffectiveDominance = territoryRules.InitialDominance
	change.NewDominance = territoryRules.InitialDominance
	change.ChangeType = HexChangeTypeOwnershipTransfer

	return change
}
