package gameengine

import (
	"testing"

	"github.com/serkan-can-eyvaz/pet-territory-wars/backend/internal/domain/id"
)

func TestRuleSetStoresAllRuleGroups(t *testing.T) {
	version, err := id.NewRuleSetVersion("mvp-0.1.0")
	if err != nil {
		t.Fatalf("create rule set version: %v", err)
	}

	rules := RuleSet{
		Version: version,
		Walk: WalkRules{
			MinDurationSeconds: 300,
			MinDistanceMeters:  300,
			MinValidRouteRatio: 0.70,
			MaxSpeedMPS:        4.5,
			MaxAccuracyMeters:  40,
			MaxJumpMeters:      250,
		},
		Movement: MovementRules{
			H3Resolution:          9,
			MinHexPresenceSeconds: 8,
			MinHexDistanceMeters:  15,
			InterpolationMeters:   10,
		},
		Territory: TerritoryRules{
			MaxDominance:          100,
			InitialDominance:      60,
			OwnerVisitGain:        10,
			EnemyAttackDamage:     15,
			DailyDecay:            5,
			MinimumOwnedDominance: 1,
			ThreatThreshold:       30,
		},
	}

	if rules.Version != version {
		t.Errorf("Version = %q, want %q", rules.Version, version)
	}
	if rules.Walk.MinDurationSeconds != 300 || rules.Walk.MaxJumpMeters != 250 {
		t.Errorf("Walk = %+v, want configured values", rules.Walk)
	}
	if rules.Movement.H3Resolution != 9 || rules.Movement.InterpolationMeters != 10 {
		t.Errorf("Movement = %+v, want configured values", rules.Movement)
	}
	if rules.Territory.MaxDominance != 100 || rules.Territory.ThreatThreshold != 30 {
		t.Errorf("Territory = %+v, want configured values", rules.Territory)
	}
}

func TestRuleSetIsComparable(t *testing.T) {
	version, err := id.NewRuleSetVersion("mvp-0.1.0")
	if err != nil {
		t.Fatalf("create rule set version: %v", err)
	}

	rules := RuleSet{Version: version}
	seen := map[RuleSet]struct{}{rules: {}}

	if _, ok := seen[rules]; !ok {
		t.Fatal("RuleSet must be comparable")
	}
}
