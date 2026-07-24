package gameengine

// resolveOwnerDefense calculates the territory change for a qualified owner visit.
func resolveOwnerDefense(
	input ResolveWalkInput,
	state HexState,
	visit VisitedHex,
	movementRules MovementRules,
	territoryRules TerritoryRules,
	effectiveDominance int,
) HexChange {
	change := HexChange{
		HexID:              state.HexID,
		ExpectedVersion:    state.Version,
		PreviousOwnerID:    state.OwnerID,
		NewOwnerID:         state.OwnerID,
		StoredDominance:    state.Dominance,
		EffectiveDominance: effectiveDominance,
		NewDominance:       effectiveDominance,
		ChangeType:         HexChangeTypeNoChange,
	}

	qualified := visit.PresenceSeconds >= movementRules.MinHexPresenceSeconds ||
		visit.DistanceMeters >= movementRules.MinHexDistanceMeters
	if state.OwnerID == nil || *state.OwnerID != input.PlayerID || !qualified {
		return change
	}

	newDominance := effectiveDominance + territoryRules.OwnerVisitGain
	if newDominance > territoryRules.MaxDominance {
		newDominance = territoryRules.MaxDominance
	}
	if newDominance == effectiveDominance {
		return change
	}

	change.EffectiveDominance = newDominance
	change.NewDominance = newDominance
	change.ChangeType = HexChangeTypeOwnerDefense

	return change
}
