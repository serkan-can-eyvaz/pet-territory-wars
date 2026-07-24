package gameengine

import (
	"fmt"
	"math"
	"strings"
)

// Validate verifies that a RuleSet can be used by the game engine.
func (rules RuleSet) Validate() error {
	if strings.TrimSpace(rules.Version.String()) == "" {
		return fmt.Errorf("RuleSet.Version must not be blank")
	}

	if err := rules.Walk.validate(); err != nil {
		return err
	}
	if err := rules.Movement.validate(); err != nil {
		return err
	}
	if err := rules.Territory.validate(); err != nil {
		return err
	}

	return nil
}

func (rules WalkRules) validate() error {
	if rules.MinDurationSeconds < 0 {
		return fmt.Errorf("RuleSet.Walk.MinDurationSeconds must be greater than or equal to zero")
	}
	if err := validateNonNegativeFinite("RuleSet.Walk.MinDistanceMeters", rules.MinDistanceMeters); err != nil {
		return err
	}
	if err := validateFinite("RuleSet.Walk.MinValidRouteRatio", rules.MinValidRouteRatio); err != nil {
		return err
	}
	if rules.MinValidRouteRatio < 0 || rules.MinValidRouteRatio > 1 {
		return fmt.Errorf("RuleSet.Walk.MinValidRouteRatio must be between zero and one")
	}
	if err := validatePositiveFinite("RuleSet.Walk.MaxSpeedMPS", rules.MaxSpeedMPS); err != nil {
		return err
	}
	if err := validateNonNegativeFinite("RuleSet.Walk.MaxAccuracyMeters", rules.MaxAccuracyMeters); err != nil {
		return err
	}
	if err := validatePositiveFinite("RuleSet.Walk.MaxJumpMeters", rules.MaxJumpMeters); err != nil {
		return err
	}

	return nil
}

func (rules MovementRules) validate() error {
	if rules.H3Resolution <= 0 {
		return fmt.Errorf("RuleSet.Movement.H3Resolution must be greater than zero")
	}
	if rules.MinHexPresenceSeconds < 0 {
		return fmt.Errorf("RuleSet.Movement.MinHexPresenceSeconds must be greater than or equal to zero")
	}
	if err := validateNonNegativeFinite("RuleSet.Movement.MinHexDistanceMeters", rules.MinHexDistanceMeters); err != nil {
		return err
	}
	if err := validatePositiveFinite("RuleSet.Movement.InterpolationMeters", rules.InterpolationMeters); err != nil {
		return err
	}

	return nil
}

func (rules TerritoryRules) validate() error {
	if rules.MaxDominance <= 0 {
		return fmt.Errorf("RuleSet.Territory.MaxDominance must be greater than zero")
	}
	if rules.InitialDominance < 1 || rules.InitialDominance > rules.MaxDominance {
		return fmt.Errorf("RuleSet.Territory.InitialDominance must be between one and RuleSet.Territory.MaxDominance")
	}
	if rules.OwnerVisitGain < 0 {
		return fmt.Errorf("RuleSet.Territory.OwnerVisitGain must be greater than or equal to zero")
	}
	if rules.EnemyAttackDamage <= 0 {
		return fmt.Errorf("RuleSet.Territory.EnemyAttackDamage must be greater than zero")
	}
	if rules.DailyDecay < 0 {
		return fmt.Errorf("RuleSet.Territory.DailyDecay must be greater than or equal to zero")
	}
	if rules.MinimumOwnedDominance < 1 || rules.MinimumOwnedDominance > rules.MaxDominance {
		return fmt.Errorf("RuleSet.Territory.MinimumOwnedDominance must be between one and RuleSet.Territory.MaxDominance")
	}
	if rules.ThreatThreshold < 0 || rules.ThreatThreshold > rules.MaxDominance {
		return fmt.Errorf("RuleSet.Territory.ThreatThreshold must be between zero and RuleSet.Territory.MaxDominance")
	}

	return nil
}

func validateNonNegativeFinite(field string, value float64) error {
	if err := validateFinite(field, value); err != nil {
		return err
	}
	if value < 0 {
		return fmt.Errorf("%s must be greater than or equal to zero", field)
	}

	return nil
}

func validatePositiveFinite(field string, value float64) error {
	if err := validateFinite(field, value); err != nil {
		return err
	}
	if value <= 0 {
		return fmt.Errorf("%s must be greater than zero", field)
	}

	return nil
}

func validateFinite(field string, value float64) error {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return fmt.Errorf("%s must be finite", field)
	}

	return nil
}
