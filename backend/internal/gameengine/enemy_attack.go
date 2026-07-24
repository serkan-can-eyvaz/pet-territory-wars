package gameengine

// resolveEnemyAttack calculates the territory change for a qualified enemy visit.
func resolveEnemyAttack(
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
	if state.OwnerID == nil || *state.OwnerID == input.PlayerID || !qualified {
		return change
	}

	afterDamage := effectiveDominance - territoryRules.EnemyAttackDamage
	if afterDamage < 0 {
		afterDamage = 0
	}
	if afterDamage == effectiveDominance {
		return change
	}

	change.EffectiveDominance = afterDamage
	change.NewDominance = afterDamage
	change.ChangeType = HexChangeTypeEnemyAttack

	return change
}
