package gameengine

// resolveEmptyCapture calculates the territory change for an unowned hex visit.
func resolveEmptyCapture(
	input ResolveWalkInput,
	state HexState,
	visit VisitedHex,
	movementRules MovementRules,
	territoryRules TerritoryRules,
) HexChange {
	change := HexChange{
		HexID:              state.HexID,
		ExpectedVersion:    state.Version,
		PreviousOwnerID:    state.OwnerID,
		NewOwnerID:         state.OwnerID,
		StoredDominance:    state.Dominance,
		EffectiveDominance: state.Dominance,
		NewDominance:       state.Dominance,
		ChangeType:         HexChangeTypeNoChange,
	}

	qualified := visit.PresenceSeconds >= movementRules.MinHexPresenceSeconds ||
		visit.DistanceMeters >= movementRules.MinHexDistanceMeters
	if state.OwnerID != nil || !qualified {
		return change
	}

	newOwnerID := input.PlayerID
	change.NewOwnerID = &newOwnerID
	change.NewDominance = territoryRules.InitialDominance
	change.ChangeType = HexChangeTypeEmptyCapture

	return change
}
