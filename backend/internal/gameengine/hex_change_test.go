package gameengine

import (
	"testing"

	"github.com/serkan-can-eyvaz/pet-territory-wars/backend/internal/domain/id"
)

func TestHexChangeStoresAllValues(t *testing.T) {
	previousOwnerID, err := id.NewPlayerID("7b0f04af-37dd-4a69-9bb1-3d8f0c59c8f8")
	if err != nil {
		t.Fatalf("create previous owner ID: %v", err)
	}
	newOwnerID, err := id.NewPlayerID("c55e1241-73df-46f2-b10d-c7b56e967024")
	if err != nil {
		t.Fatalf("create new owner ID: %v", err)
	}
	hexID := id.NewHexID(0x8928308280fffff)
	change := HexChange{
		HexID:              hexID,
		ExpectedVersion:    12,
		PreviousOwnerID:    &previousOwnerID,
		NewOwnerID:         &newOwnerID,
		StoredDominance:    80,
		EffectiveDominance: 75,
		NewDominance:       60,
		ChangeType:         HexChangeTypeEnemyAttack,
		Reason:             HexChangeReason("attack_applied"),
	}

	if change.HexID != hexID || change.ExpectedVersion != 12 {
		t.Error("HexChange did not preserve HexID or ExpectedVersion")
	}
	if change.PreviousOwnerID == nil || *change.PreviousOwnerID != previousOwnerID {
		t.Error("HexChange did not preserve PreviousOwnerID")
	}
	if change.NewOwnerID == nil || *change.NewOwnerID != newOwnerID {
		t.Error("HexChange did not preserve NewOwnerID")
	}
	if change.StoredDominance != 80 || change.EffectiveDominance != 75 || change.NewDominance != 60 {
		t.Errorf("HexChange dominance values = %+v, want configured values", change)
	}
	if change.ChangeType != HexChangeTypeEnemyAttack {
		t.Errorf("ChangeType = %q, want %q", change.ChangeType, HexChangeTypeEnemyAttack)
	}
	if change.Reason != HexChangeReason("attack_applied") {
		t.Errorf("Reason = %q, want %q", change.Reason, HexChangeReason("attack_applied"))
	}
}

func TestHexChangeSupportsUnownedOwners(t *testing.T) {
	change := HexChange{
		HexID:      id.NewHexID(0x8928308280fffff),
		ChangeType: HexChangeTypeNoChange,
	}

	if change.PreviousOwnerID != nil {
		t.Errorf("PreviousOwnerID = %v, want nil", change.PreviousOwnerID)
	}
	if change.NewOwnerID != nil {
		t.Errorf("NewOwnerID = %v, want nil", change.NewOwnerID)
	}
}

func TestHexChangeTypeCanonicalValues(t *testing.T) {
	tests := []struct {
		name  string
		value HexChangeType
		want  string
	}{
		{name: "empty capture", value: HexChangeTypeEmptyCapture, want: "EMPTY_CAPTURE"},
		{name: "owner defense", value: HexChangeTypeOwnerDefense, want: "OWNER_DEFENSE"},
		{name: "enemy attack", value: HexChangeTypeEnemyAttack, want: "ENEMY_ATTACK"},
		{name: "ownership transfer", value: HexChangeTypeOwnershipTransfer, want: "OWNERSHIP_TRANSFER"},
		{name: "no change", value: HexChangeTypeNoChange, want: "NO_CHANGE"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if string(test.value) != test.want {
				t.Errorf("HexChangeType = %q, want %q", test.value, test.want)
			}
		})
	}
}
