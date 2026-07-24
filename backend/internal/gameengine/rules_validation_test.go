package gameengine

import (
	"math"
	"strings"
	"testing"

	"github.com/serkan-can-eyvaz/pet-territory-wars/backend/internal/domain/id"
)

func TestRuleSetValidateAcceptsBaseline(t *testing.T) {
	if err := validRuleSet(t).Validate(); err != nil {
		t.Fatalf("validate RuleSet: %v", err)
	}
}

func TestRuleSetValidateRejectsInvalidValues(t *testing.T) {
	tests := []struct {
		name          string
		modify        func(*RuleSet)
		expectedField string
	}{
		{
			name: "blank version",
			modify: func(rules *RuleSet) {
				rules.Version = ""
			},
			expectedField: "RuleSet.Version",
		},
		{
			name: "whitespace version",
			modify: func(rules *RuleSet) {
				rules.Version = " \t\n "
			},
			expectedField: "RuleSet.Version",
		},
		{
			name: "negative minimum duration",
			modify: func(rules *RuleSet) {
				rules.Walk.MinDurationSeconds = -1
			},
			expectedField: "RuleSet.Walk.MinDurationSeconds",
		},
		{
			name: "negative minimum distance",
			modify: func(rules *RuleSet) {
				rules.Walk.MinDistanceMeters = -1
			},
			expectedField: "RuleSet.Walk.MinDistanceMeters",
		},
		{
			name: "route ratio below zero",
			modify: func(rules *RuleSet) {
				rules.Walk.MinValidRouteRatio = -0.01
			},
			expectedField: "RuleSet.Walk.MinValidRouteRatio",
		},
		{
			name: "route ratio above one",
			modify: func(rules *RuleSet) {
				rules.Walk.MinValidRouteRatio = 1.01
			},
			expectedField: "RuleSet.Walk.MinValidRouteRatio",
		},
		{
			name: "non-positive maximum speed",
			modify: func(rules *RuleSet) {
				rules.Walk.MaxSpeedMPS = 0
			},
			expectedField: "RuleSet.Walk.MaxSpeedMPS",
		},
		{
			name: "negative maximum accuracy",
			modify: func(rules *RuleSet) {
				rules.Walk.MaxAccuracyMeters = -1
			},
			expectedField: "RuleSet.Walk.MaxAccuracyMeters",
		},
		{
			name: "non-positive maximum jump",
			modify: func(rules *RuleSet) {
				rules.Walk.MaxJumpMeters = 0
			},
			expectedField: "RuleSet.Walk.MaxJumpMeters",
		},
		{
			name: "non-positive H3 resolution",
			modify: func(rules *RuleSet) {
				rules.Movement.H3Resolution = 0
			},
			expectedField: "RuleSet.Movement.H3Resolution",
		},
		{
			name: "negative hex presence",
			modify: func(rules *RuleSet) {
				rules.Movement.MinHexPresenceSeconds = -1
			},
			expectedField: "RuleSet.Movement.MinHexPresenceSeconds",
		},
		{
			name: "negative hex distance",
			modify: func(rules *RuleSet) {
				rules.Movement.MinHexDistanceMeters = -1
			},
			expectedField: "RuleSet.Movement.MinHexDistanceMeters",
		},
		{
			name: "non-positive interpolation distance",
			modify: func(rules *RuleSet) {
				rules.Movement.InterpolationMeters = 0
			},
			expectedField: "RuleSet.Movement.InterpolationMeters",
		},
		{
			name: "non-positive maximum dominance",
			modify: func(rules *RuleSet) {
				rules.Territory.MaxDominance = 0
			},
			expectedField: "RuleSet.Territory.MaxDominance",
		},
		{
			name: "initial dominance below one",
			modify: func(rules *RuleSet) {
				rules.Territory.InitialDominance = 0
			},
			expectedField: "RuleSet.Territory.InitialDominance",
		},
		{
			name: "initial dominance above maximum",
			modify: func(rules *RuleSet) {
				rules.Territory.InitialDominance = rules.Territory.MaxDominance + 1
			},
			expectedField: "RuleSet.Territory.InitialDominance",
		},
		{
			name: "negative owner visit gain",
			modify: func(rules *RuleSet) {
				rules.Territory.OwnerVisitGain = -1
			},
			expectedField: "RuleSet.Territory.OwnerVisitGain",
		},
		{
			name: "non-positive enemy attack damage",
			modify: func(rules *RuleSet) {
				rules.Territory.EnemyAttackDamage = 0
			},
			expectedField: "RuleSet.Territory.EnemyAttackDamage",
		},
		{
			name: "negative daily decay",
			modify: func(rules *RuleSet) {
				rules.Territory.DailyDecay = -1
			},
			expectedField: "RuleSet.Territory.DailyDecay",
		},
		{
			name: "minimum owned dominance below one",
			modify: func(rules *RuleSet) {
				rules.Territory.MinimumOwnedDominance = 0
			},
			expectedField: "RuleSet.Territory.MinimumOwnedDominance",
		},
		{
			name: "minimum owned dominance above maximum",
			modify: func(rules *RuleSet) {
				rules.Territory.MinimumOwnedDominance = rules.Territory.MaxDominance + 1
			},
			expectedField: "RuleSet.Territory.MinimumOwnedDominance",
		},
		{
			name: "negative threat threshold",
			modify: func(rules *RuleSet) {
				rules.Territory.ThreatThreshold = -1
			},
			expectedField: "RuleSet.Territory.ThreatThreshold",
		},
		{
			name: "threat threshold above maximum",
			modify: func(rules *RuleSet) {
				rules.Territory.ThreatThreshold = rules.Territory.MaxDominance + 1
			},
			expectedField: "RuleSet.Territory.ThreatThreshold",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rules := validRuleSet(t)
			test.modify(&rules)

			err := rules.Validate()
			if err == nil {
				t.Fatal("expected validation error")
			}
			if !strings.Contains(err.Error(), test.expectedField) {
				t.Errorf("error = %q, want field %q", err, test.expectedField)
			}
		})
	}
}

func TestRuleSetValidateRejectsNonFiniteFloats(t *testing.T) {
	values := []struct {
		name  string
		value float64
	}{
		{name: "NaN", value: math.NaN()},
		{name: "positive infinity", value: math.Inf(1)},
		{name: "negative infinity", value: math.Inf(-1)},
	}
	fields := []struct {
		name          string
		set           func(*RuleSet, float64)
		expectedField string
	}{
		{
			name: "minimum distance",
			set: func(rules *RuleSet, value float64) {
				rules.Walk.MinDistanceMeters = value
			},
			expectedField: "RuleSet.Walk.MinDistanceMeters",
		},
		{
			name: "minimum valid route ratio",
			set: func(rules *RuleSet, value float64) {
				rules.Walk.MinValidRouteRatio = value
			},
			expectedField: "RuleSet.Walk.MinValidRouteRatio",
		},
		{
			name: "maximum speed",
			set: func(rules *RuleSet, value float64) {
				rules.Walk.MaxSpeedMPS = value
			},
			expectedField: "RuleSet.Walk.MaxSpeedMPS",
		},
		{
			name: "maximum accuracy",
			set: func(rules *RuleSet, value float64) {
				rules.Walk.MaxAccuracyMeters = value
			},
			expectedField: "RuleSet.Walk.MaxAccuracyMeters",
		},
		{
			name: "maximum jump",
			set: func(rules *RuleSet, value float64) {
				rules.Walk.MaxJumpMeters = value
			},
			expectedField: "RuleSet.Walk.MaxJumpMeters",
		},
		{
			name: "minimum hex distance",
			set: func(rules *RuleSet, value float64) {
				rules.Movement.MinHexDistanceMeters = value
			},
			expectedField: "RuleSet.Movement.MinHexDistanceMeters",
		},
		{
			name: "interpolation distance",
			set: func(rules *RuleSet, value float64) {
				rules.Movement.InterpolationMeters = value
			},
			expectedField: "RuleSet.Movement.InterpolationMeters",
		},
	}

	for _, field := range fields {
		for _, value := range values {
			t.Run(field.name+"/"+value.name, func(t *testing.T) {
				rules := validRuleSet(t)
				field.set(&rules, value.value)

				err := rules.Validate()
				if err == nil {
					t.Fatal("expected validation error")
				}
				if !strings.Contains(err.Error(), field.expectedField) {
					t.Errorf("error = %q, want field %q", err, field.expectedField)
				}
			})
		}
	}
}

func TestRuleSetValidateReturnsFirstError(t *testing.T) {
	tests := []struct {
		name          string
		modify        func(*RuleSet)
		expectedField string
	}{
		{
			name: "version before all rule groups",
			modify: func(rules *RuleSet) {
				rules.Version = ""
				rules.Walk.MinDurationSeconds = -1
				rules.Movement.H3Resolution = 0
				rules.Territory.MaxDominance = 0
			},
			expectedField: "RuleSet.Version",
		},
		{
			name: "walk before movement and territory",
			modify: func(rules *RuleSet) {
				rules.Walk.MinDurationSeconds = -1
				rules.Movement.H3Resolution = 0
				rules.Territory.MaxDominance = 0
			},
			expectedField: "RuleSet.Walk.MinDurationSeconds",
		},
		{
			name: "movement before territory",
			modify: func(rules *RuleSet) {
				rules.Movement.H3Resolution = 0
				rules.Territory.MaxDominance = 0
			},
			expectedField: "RuleSet.Movement.H3Resolution",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rules := validRuleSet(t)
			test.modify(&rules)

			err := rules.Validate()
			if err == nil {
				t.Fatal("expected validation error")
			}
			if !strings.Contains(err.Error(), test.expectedField) {
				t.Errorf("error = %q, want first field %q", err, test.expectedField)
			}
		})
	}
}

func validRuleSet(t *testing.T) RuleSet {
	t.Helper()

	version, err := id.NewRuleSetVersion("mvp-0.1.0")
	if err != nil {
		t.Fatalf("create rule set version: %v", err)
	}

	return RuleSet{
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
}
