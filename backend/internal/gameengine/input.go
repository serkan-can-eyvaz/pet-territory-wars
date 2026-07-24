package gameengine

import (
	"time"

	"github.com/serkan-can-eyvaz/pet-territory-wars/backend/internal/domain/id"
)

// ResolveWalkInput contains all state required to resolve a walk.
type ResolveWalkInput struct {
	WalkID        id.WalkID
	PlayerID      id.PlayerID
	CityID        id.CityID
	StartedAt     time.Time
	EndedAt       time.Time
	EvaluatedAt   time.Time
	Route         []LocationPoint
	ExistingHexes map[id.HexID]HexState
	Boundary      CityBoundary
	EngineVersion id.EngineVersion
}

// LocationPoint represents a recorded location in a walk route.
type LocationPoint struct {
	Sequence       int
	Latitude       float64
	Longitude      float64
	RecordedAt     time.Time
	AccuracyMeters float64
	IsMockLocation bool
}

// HexState represents the current state of a territory hexagon.
type HexState struct {
	HexID         id.HexID
	OwnerID       *id.PlayerID
	Dominance     int
	LastUpdatedAt time.Time
	Version       int64
}
