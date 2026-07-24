package gameengine

import "time"

func calculateEffectiveDominance(hexState HexState, rules TerritoryRules, evaluatedAt time.Time) int {
	if hexState.OwnerID == nil {
		return 0
	}
	if rules.DailyDecay == 0 {
		return hexState.Dominance
	}
	if !evaluatedAt.After(hexState.LastUpdatedAt) {
		return hexState.Dominance
	}

	elapsedFullDays := int(evaluatedAt.Sub(hexState.LastUpdatedAt) / (24 * time.Hour))
	decayAmount := elapsedFullDays * rules.DailyDecay
	effectiveDominance := hexState.Dominance - decayAmount
	if effectiveDominance < rules.MinimumOwnedDominance {
		return rules.MinimumOwnedDominance
	}

	return effectiveDominance
}
