// Package gameengine contains the deterministic game calculation model.
package gameengine

import "github.com/serkan-can-eyvaz/pet-territory-wars/backend/internal/domain/id"

// RuleSet contains the immutable, versioned configuration consumed by the game engine.
type RuleSet struct {
	Version   id.RuleSetVersion
	Walk      WalkRules
	Movement  MovementRules
	Territory TerritoryRules
}

// WalkRules contains rules for walk validity and segment evaluation.
type WalkRules struct {
	MinDurationSeconds int
	MinDistanceMeters  float64
	MinValidRouteRatio float64
	MaxSpeedMPS        float64
	MaxAccuracyMeters  float64
	MaxJumpMeters      float64
}

// MovementRules contains rules for route interpolation and hex visits.
type MovementRules struct {
	H3Resolution          int
	MinHexPresenceSeconds int
	MinHexDistanceMeters  float64
	InterpolationMeters   float64
}

// TerritoryRules contains rules for territory dominance changes.
type TerritoryRules struct {
	MaxDominance          int
	InitialDominance      int
	OwnerVisitGain        int
	EnemyAttackDamage     int
	DailyDecay            int
	MinimumOwnedDominance int
	ThreatThreshold       int
}
