package gameengine

import "github.com/serkan-can-eyvaz/pet-territory-wars/backend/internal/domain/id"

// HexChangeType identifies the type of territory change produced for a hex.
type HexChangeType string

const (
	HexChangeTypeEmptyCapture      HexChangeType = "EMPTY_CAPTURE"
	HexChangeTypeOwnerDefense      HexChangeType = "OWNER_DEFENSE"
	HexChangeTypeEnemyAttack       HexChangeType = "ENEMY_ATTACK"
	HexChangeTypeOwnershipTransfer HexChangeType = "OWNERSHIP_TRANSFER"
	HexChangeTypeNoChange          HexChangeType = "NO_CHANGE"
)

// HexChangeReason is a deterministic explanation code for a hex change.
type HexChangeReason string

// HexChange describes the resulting state transition for a territory hexagon.
type HexChange struct {
	HexID              id.HexID
	ExpectedVersion    int64
	PreviousOwnerID    *id.PlayerID
	NewOwnerID         *id.PlayerID
	StoredDominance    int
	EffectiveDominance int
	NewDominance       int
	ChangeType         HexChangeType
	Reason             HexChangeReason
}
