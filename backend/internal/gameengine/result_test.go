package gameengine

import (
	"testing"
	"time"

	"github.com/serkan-can-eyvaz/pet-territory-wars/backend/internal/domain/id"
)

func TestResultCanonicalConstants(t *testing.T) {
	if string(WalkResolutionStatusResolved) != "RESOLVED" || string(WalkResolutionStatusRejected) != "REJECTED" {
		t.Error("WalkResolutionStatus constants do not match the contract")
	}
	if string(WalkValidationStatusValid) != "VALID" || string(WalkValidationStatusPartiallyValid) != "PARTIALLY_VALID" || string(WalkValidationStatusInvalid) != "INVALID" {
		t.Error("WalkValidationStatus constants do not match the contract")
	}
	if string(WalkRejectionReasonTooShortDuration) != "TOO_SHORT_DURATION" ||
		string(WalkRejectionReasonTooShortDistance) != "TOO_SHORT_DISTANCE" ||
		string(WalkRejectionReasonLowValidRouteRatio) != "LOW_VALID_ROUTE_RATIO" ||
		string(WalkRejectionReasonNoValidSegments) != "NO_VALID_SEGMENTS" ||
		string(WalkRejectionReasonMockLocationDetected) != "MOCK_LOCATION_DETECTED" ||
		string(WalkRejectionReasonOutsideActiveCity) != "OUTSIDE_ACTIVE_CITY" {
		t.Error("WalkRejectionReason constants do not match the contract")
	}
	if string(ScoreMetricActiveHex) != "ACTIVE_HEX" ||
		string(ScoreMetricCapture) != "CAPTURE" ||
		string(ScoreMetricLifetimeCapture) != "LIFETIME_CAPTURE" ||
		string(ScoreMetricDefense) != "DEFENSE" ||
		string(ScoreMetricSteal) != "STEAL" ||
		string(ScoreMetricLifetimeSteal) != "LIFETIME_STEAL" {
		t.Error("ScoreMetric constants do not match the contract")
	}
	if string(EventTypeWalkValidated) != "WALK_VALIDATED" ||
		string(EventTypeWalkPartiallyValidated) != "WALK_PARTIALLY_VALIDATED" ||
		string(EventTypeWalkRejected) != "WALK_REJECTED" ||
		string(EventTypeEmptyHexCaptured) != "EMPTY_HEX_CAPTURED" ||
		string(EventTypeHexDefended) != "HEX_DEFENDED" ||
		string(EventTypeHexAttacked) != "HEX_ATTACKED" ||
		string(EventTypeHexUnderThreat) != "HEX_UNDER_THREAT" ||
		string(EventTypeHexOwnershipTransferred) != "HEX_OWNERSHIP_TRANSFERRED" ||
		string(EventTypePlayerScoreChanged) != "PLAYER_SCORE_CHANGED" {
		t.Error("EventType constants do not match the contract")
	}
}

func TestResultModelsStoreData(t *testing.T) {
	walkID, err := id.NewWalkID("7b0f04af-37dd-4a69-9bb1-3d8f0c59c8f8")
	if err != nil {
		t.Fatalf("create walk ID: %v", err)
	}
	playerID, err := id.NewPlayerID("c55e1241-73df-46f2-b10d-c7b56e967024")
	if err != nil {
		t.Fatalf("create player ID: %v", err)
	}
	eventID, err := id.NewEventID("1276c9f1-97a9-48a8-81d2-063d80a221d5")
	if err != nil {
		t.Fatalf("create event ID: %v", err)
	}
	ruleSetVersion, err := id.NewRuleSetVersion("mvp-0.1.0")
	if err != nil {
		t.Fatalf("create rule set version: %v", err)
	}
	engineVersion, err := id.NewEngineVersion("1.0.0")
	if err != nil {
		t.Fatalf("create engine version: %v", err)
	}
	hexID := id.NewHexID(0x8928308280fffff)
	firstEnteredAt := time.Date(2026, time.July, 24, 10, 0, 0, 0, time.UTC)
	lastExitedAt := firstEnteredAt.Add(2 * time.Minute)
	evaluatedAt := lastExitedAt.Add(time.Minute)
	validation := ValidationResult{
		Status:           WalkValidationStatusInvalid,
		RejectionReasons: []WalkRejectionReason{WalkRejectionReasonTooShortDuration, WalkRejectionReasonTooShortDistance},
	}
	emptyValidation := ValidationResult{Status: WalkValidationStatusValid, RejectionReasons: []WalkRejectionReason{}}
	metrics := WalkMetrics{
		TotalDistanceMeters:  450.5,
		ValidDistanceMeters:  320.25,
		ValidDurationSeconds: 300,
		ValidRouteRatio:      0.71,
	}
	visitedHex := VisitedHex{
		HexID:           hexID,
		PresenceSeconds: 42,
		DistanceMeters:  125.75,
		EntryCount:      2,
		FirstEnteredAt:  firstEnteredAt,
		LastExitedAt:    lastExitedAt,
	}
	hexChange := HexChange{
		HexID:              hexID,
		ExpectedVersion:    4,
		StoredDominance:    80,
		EffectiveDominance: 75,
		NewDominance:       60,
		ChangeType:         HexChangeTypeEnemyAttack,
		Reason:             HexChangeReason("attack_applied"),
	}
	scoreChanges := []ScoreChange{
		{PlayerID: playerID, Metric: ScoreMetricCapture, Delta: 1, Reason: ScoreChangeReason("empty_capture"), WalkID: walkID},
		{PlayerID: playerID, Metric: ScoreMetricActiveHex, Delta: -1, Reason: ScoreChangeReason("ownership_transfer"), WalkID: walkID},
	}
	metadata := CalculationMetadata{
		EngineVersion:  engineVersion,
		RuleSetVersion: ruleSetVersion,
		InputHash:      "abc123",
		EvaluatedAt:    evaluatedAt,
	}
	result := ResolveWalkResult{
		Status:       WalkResolutionStatusRejected,
		Validation:   validation,
		Metrics:      metrics,
		VisitedHexes: []VisitedHex{visitedHex},
		HexChanges:   []HexChange{hexChange},
		ScoreChanges: scoreChanges,
		Events: []DomainEvent{
			{
				ID:             eventID,
				Type:           EventTypeWalkRejected,
				OccurredAt:     evaluatedAt,
				WalkID:         walkID,
				PlayerID:       playerID,
				HexID:          nil,
				RuleSetVersion: ruleSetVersion,
				EngineVersion:  engineVersion,
				Payload:        nil,
			},
			{
				ID:             eventID,
				Type:           EventTypeHexAttacked,
				OccurredAt:     evaluatedAt,
				WalkID:         walkID,
				PlayerID:       playerID,
				HexID:          &hexID,
				RuleSetVersion: ruleSetVersion,
				EngineVersion:  engineVersion,
				Payload:        struct{ Description string }{Description: "attack"},
			},
		},
		Metadata: metadata,
	}

	if len(emptyValidation.RejectionReasons) != 0 {
		t.Error("ValidationResult did not preserve an empty rejection list")
	}
	if result.Validation.Status != validation.Status || len(result.Validation.RejectionReasons) != 2 ||
		result.Validation.RejectionReasons[0] != WalkRejectionReasonTooShortDuration ||
		result.Validation.RejectionReasons[1] != WalkRejectionReasonTooShortDistance {
		t.Error("ValidationResult did not preserve ordered rejection reasons")
	}
	if result.Metrics != metrics {
		t.Error("WalkMetrics did not preserve values")
	}
	if len(result.VisitedHexes) != 1 || result.VisitedHexes[0] != visitedHex {
		t.Error("VisitedHexes did not preserve values")
	}
	if len(result.HexChanges) != 1 || result.HexChanges[0] != hexChange {
		t.Error("HexChanges did not reuse the existing HexChange value")
	}
	if len(result.ScoreChanges) != 2 || result.ScoreChanges[0] != scoreChanges[0] || result.ScoreChanges[1] != scoreChanges[1] {
		t.Error("ScoreChanges did not preserve positive and negative deltas")
	}
	if len(result.Events) != 2 || result.Events[0].HexID != nil || result.Events[0].Payload != nil {
		t.Error("DomainEvent did not preserve the walk-level event contract")
	}
	if result.Events[1].HexID == nil || *result.Events[1].HexID != hexID || result.Events[1].Payload == nil {
		t.Error("DomainEvent did not preserve the hex event contract")
	}
	if result.Metadata != metadata {
		t.Error("CalculationMetadata did not preserve values")
	}
	if result.Status != WalkResolutionStatusRejected {
		t.Errorf("Status = %q, want %q", result.Status, WalkResolutionStatusRejected)
	}
}
