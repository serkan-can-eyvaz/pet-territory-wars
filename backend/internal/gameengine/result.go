package gameengine

import (
	"time"

	"github.com/serkan-can-eyvaz/pet-territory-wars/backend/internal/domain/id"
)

// WalkResolutionStatus identifies the outcome of engine resolution.
type WalkResolutionStatus string

const (
	WalkResolutionStatusResolved WalkResolutionStatus = "RESOLVED"
	WalkResolutionStatusRejected WalkResolutionStatus = "REJECTED"
)

// WalkValidationStatus identifies the business validity of a walk.
type WalkValidationStatus string

const (
	WalkValidationStatusValid          WalkValidationStatus = "VALID"
	WalkValidationStatusPartiallyValid WalkValidationStatus = "PARTIALLY_VALID"
	WalkValidationStatusInvalid        WalkValidationStatus = "INVALID"
)

// WalkRejectionReason identifies a business reason for rejecting a walk.
type WalkRejectionReason string

const (
	WalkRejectionReasonTooShortDuration     WalkRejectionReason = "TOO_SHORT_DURATION"
	WalkRejectionReasonTooShortDistance     WalkRejectionReason = "TOO_SHORT_DISTANCE"
	WalkRejectionReasonLowValidRouteRatio   WalkRejectionReason = "LOW_VALID_ROUTE_RATIO"
	WalkRejectionReasonNoValidSegments      WalkRejectionReason = "NO_VALID_SEGMENTS"
	WalkRejectionReasonMockLocationDetected WalkRejectionReason = "MOCK_LOCATION_DETECTED"
	WalkRejectionReasonOutsideActiveCity    WalkRejectionReason = "OUTSIDE_ACTIVE_CITY"
)

// ValidationResult carries the business validity outcome of a walk.
type ValidationResult struct {
	Status           WalkValidationStatus
	RejectionReasons []WalkRejectionReason
}

// WalkMetrics carries measured distances, duration, and valid route ratio.
type WalkMetrics struct {
	TotalDistanceMeters  float64
	ValidDistanceMeters  float64
	ValidDurationSeconds int
	ValidRouteRatio      float64
}

// VisitedHex carries the aggregated visit data for a hexagon.
type VisitedHex struct {
	HexID           id.HexID
	PresenceSeconds int
	DistanceMeters  float64
	EntryCount      int
	FirstEnteredAt  time.Time
	LastExitedAt    time.Time
}

// ScoreMetric identifies a player score dimension.
type ScoreMetric string

const (
	ScoreMetricActiveHex       ScoreMetric = "ACTIVE_HEX"
	ScoreMetricCapture         ScoreMetric = "CAPTURE"
	ScoreMetricLifetimeCapture ScoreMetric = "LIFETIME_CAPTURE"
	ScoreMetricDefense         ScoreMetric = "DEFENSE"
	ScoreMetricSteal           ScoreMetric = "STEAL"
	ScoreMetricLifetimeSteal   ScoreMetric = "LIFETIME_STEAL"
)

// ScoreChangeReason is a deterministic explanation code for a score change.
type ScoreChangeReason string

const (
	ScoreChangeReasonEmptyCapture      ScoreChangeReason = "EMPTY_CAPTURE"
	ScoreChangeReasonOwnerDefense      ScoreChangeReason = "OWNER_DEFENSE"
	ScoreChangeReasonOwnershipTransfer ScoreChangeReason = "OWNERSHIP_TRANSFER"
)

// ScoreChange carries a score delta for a player.
type ScoreChange struct {
	PlayerID id.PlayerID
	Metric   ScoreMetric
	Delta    int64
	Reason   ScoreChangeReason
	WalkID   id.WalkID
}

// EventType identifies a domain event produced by the engine.
type EventType string

const (
	EventTypeWalkValidated           EventType = "WALK_VALIDATED"
	EventTypeWalkPartiallyValidated  EventType = "WALK_PARTIALLY_VALIDATED"
	EventTypeWalkRejected            EventType = "WALK_REJECTED"
	EventTypeEmptyHexCaptured        EventType = "EMPTY_HEX_CAPTURED"
	EventTypeHexDefended             EventType = "HEX_DEFENDED"
	EventTypeHexAttacked             EventType = "HEX_ATTACKED"
	EventTypeHexUnderThreat          EventType = "HEX_UNDER_THREAT"
	EventTypeHexOwnershipTransferred EventType = "HEX_OWNERSHIP_TRANSFERRED"
	EventTypePlayerScoreChanged      EventType = "PLAYER_SCORE_CHANGED"
)

// EventPayload carries immutable event-specific value data.
type EventPayload interface{}

// DomainEvent carries the common fields of an engine domain event.
type DomainEvent struct {
	ID             id.EventID
	Type           EventType
	OccurredAt     time.Time
	WalkID         id.WalkID
	PlayerID       id.PlayerID
	HexID          *id.HexID
	RuleSetVersion id.RuleSetVersion
	EngineVersion  id.EngineVersion
	Payload        EventPayload
}

// CalculationMetadata carries the immutable metadata associated with resolution.
type CalculationMetadata struct {
	EngineVersion  id.EngineVersion
	RuleSetVersion id.RuleSetVersion
	InputHash      string
	EvaluatedAt    time.Time
}

// ResolveWalkResult carries the complete output of resolving a walk.
type ResolveWalkResult struct {
	Status       WalkResolutionStatus
	Validation   ValidationResult
	Metrics      WalkMetrics
	VisitedHexes []VisitedHex
	HexChanges   []HexChange
	ScoreChanges []ScoreChange
	Events       []DomainEvent
	Metadata     CalculationMetadata
}
