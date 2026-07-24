package gameengine

import (
	"testing"
	"time"

	"github.com/serkan-can-eyvaz/pet-territory-wars/backend/internal/domain/id"
)

func TestCalculateEffectiveDominance_UnownedHexIsZero(t *testing.T) {
	t.Parallel()

	hexState := HexState{Dominance: 50}
	evaluatedAt := time.Date(2026, time.January, 1, 12, 0, 0, 0, time.UTC)
	got := calculateEffectiveDominance(hexState, TerritoryRules{DailyDecay: 5, MinimumOwnedDominance: 1}, evaluatedAt)
	if got != 0 {
		t.Fatalf("effective dominance = %d, want 0", got)
	}
}

func TestCalculateEffectiveDominance_DailyDecayZeroReturnsStoredDominance(t *testing.T) {
	t.Parallel()

	hexState := ownedHexState(60, time.Date(2026, time.January, 1, 12, 0, 0, 0, time.UTC))
	got := calculateEffectiveDominance(hexState, TerritoryRules{DailyDecay: 0, MinimumOwnedDominance: 1}, hexState.LastUpdatedAt.Add(1000*24*time.Hour))
	if got != 60 {
		t.Fatalf("effective dominance = %d, want 60", got)
	}
}

func TestCalculateEffectiveDominance_AppliesOnlyFullDays(t *testing.T) {
	t.Parallel()

	lastUpdatedAt := time.Date(2026, time.January, 1, 10, 0, 0, 0, time.FixedZone("local", 3*60*60))
	hexState := ownedHexState(60, lastUpdatedAt)
	rules := TerritoryRules{DailyDecay: 5, MinimumOwnedDominance: 1}

	if got := calculateEffectiveDominance(hexState, rules, lastUpdatedAt.Add(2*24*time.Hour+23*time.Hour)); got != 50 {
		t.Fatalf("partial-day effective dominance = %d, want 50", got)
	}
	if got := calculateEffectiveDominance(hexState, rules, lastUpdatedAt.Add(3*24*time.Hour)); got != 45 {
		t.Fatalf("three-day effective dominance = %d, want 45", got)
	}
}

func TestCalculateEffectiveDominance_DoesNotDecayAtOrBeforeLastUpdate(t *testing.T) {
	t.Parallel()

	lastUpdatedAt := time.Date(2026, time.January, 1, 12, 0, 0, 0, time.UTC)
	hexState := ownedHexState(60, lastUpdatedAt)
	rules := TerritoryRules{DailyDecay: 5, MinimumOwnedDominance: 1}

	for _, evaluatedAt := range []time.Time{lastUpdatedAt, lastUpdatedAt.Add(-24 * time.Hour)} {
		if got := calculateEffectiveDominance(hexState, rules, evaluatedAt); got != 60 {
			t.Fatalf("evaluated at %s: effective dominance = %d, want 60", evaluatedAt, got)
		}
	}
}

func TestCalculateEffectiveDominance_StopsAtMinimumOwnedDominance(t *testing.T) {
	t.Parallel()

	hexState := ownedHexState(3, time.Date(2026, time.January, 1, 12, 0, 0, 0, time.UTC))
	got := calculateEffectiveDominance(hexState, TerritoryRules{DailyDecay: 5, MinimumOwnedDominance: 1}, hexState.LastUpdatedAt.Add(10*24*time.Hour))
	if got != 1 {
		t.Fatalf("effective dominance = %d, want 1", got)
	}
}

func TestCalculateEffectiveDominance_IsDeterministicAndDoesNotMutateInput(t *testing.T) {
	t.Parallel()

	hexState := ownedHexState(100, time.Date(2026, time.January, 1, 12, 0, 0, 0, time.UTC))
	original := hexState
	rules := TerritoryRules{MaxDominance: 100, DailyDecay: 5, MinimumOwnedDominance: 1}
	evaluatedAt := hexState.LastUpdatedAt.Add(3 * 24 * time.Hour)

	first := calculateEffectiveDominance(hexState, rules, evaluatedAt)
	second := calculateEffectiveDominance(hexState, rules, evaluatedAt)
	if first != 85 || second != first {
		t.Fatalf("results = %d and %d, want deterministic 85", first, second)
	}
	if hexState != original {
		t.Fatal("hex state was mutated")
	}
}

func ownedHexState(dominance int, lastUpdatedAt time.Time) HexState {
	var ownerID id.PlayerID
	return HexState{
		OwnerID:       &ownerID,
		Dominance:     dominance,
		LastUpdatedAt: lastUpdatedAt,
	}
}
