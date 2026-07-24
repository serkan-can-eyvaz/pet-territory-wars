package gameengine

import "github.com/serkan-can-eyvaz/pet-territory-wars/backend/internal/domain/id"

type domainEventContext struct {
	Input          ResolveWalkInput
	RuleSetVersion id.RuleSetVersion
	TerritoryRules TerritoryRules
	AttackChange   HexChange
	FinalChange    HexChange
	ScoreChanges   []ScoreChange
	EventIDs       []id.EventID
}

// resolveDomainEvents converts resolved territory and score changes into canonical domain events.
func resolveDomainEvents(context domainEventContext) []DomainEvent {
	events := make([]DomainEvent, 0)
	if context.FinalChange.ChangeType == HexChangeTypeNoChange {
		return events
	}

	switch context.FinalChange.ChangeType {
	case HexChangeTypeEmptyCapture:
		hexID := context.FinalChange.HexID
		events = append(events, DomainEvent{
			ID:             context.EventIDs[len(events)],
			Type:           EventTypeEmptyHexCaptured,
			OccurredAt:     context.Input.EvaluatedAt,
			WalkID:         context.Input.WalkID,
			PlayerID:       context.Input.PlayerID,
			HexID:          &hexID,
			RuleSetVersion: context.RuleSetVersion,
			EngineVersion:  context.Input.EngineVersion,
		})
	case HexChangeTypeOwnerDefense:
		hexID := context.FinalChange.HexID
		events = append(events, DomainEvent{
			ID:             context.EventIDs[len(events)],
			Type:           EventTypeHexDefended,
			OccurredAt:     context.Input.EvaluatedAt,
			WalkID:         context.Input.WalkID,
			PlayerID:       context.Input.PlayerID,
			HexID:          &hexID,
			RuleSetVersion: context.RuleSetVersion,
			EngineVersion:  context.Input.EngineVersion,
		})
	case HexChangeTypeEnemyAttack, HexChangeTypeOwnershipTransfer:
		attackHexID := context.AttackChange.HexID
		events = append(events, DomainEvent{
			ID:             context.EventIDs[len(events)],
			Type:           EventTypeHexAttacked,
			OccurredAt:     context.Input.EvaluatedAt,
			WalkID:         context.Input.WalkID,
			PlayerID:       context.Input.PlayerID,
			HexID:          &attackHexID,
			RuleSetVersion: context.RuleSetVersion,
			EngineVersion:  context.Input.EngineVersion,
		})

		if context.AttackChange.NewDominance > 0 && context.AttackChange.NewDominance <= context.TerritoryRules.ThreatThreshold {
			events = append(events, DomainEvent{
				ID:             context.EventIDs[len(events)],
				Type:           EventTypeHexUnderThreat,
				OccurredAt:     context.Input.EvaluatedAt,
				WalkID:         context.Input.WalkID,
				PlayerID:       context.Input.PlayerID,
				HexID:          &attackHexID,
				RuleSetVersion: context.RuleSetVersion,
				EngineVersion:  context.Input.EngineVersion,
			})
		}

		if context.FinalChange.ChangeType == HexChangeTypeOwnershipTransfer {
			finalHexID := context.FinalChange.HexID
			events = append(events, DomainEvent{
				ID:             context.EventIDs[len(events)],
				Type:           EventTypeHexOwnershipTransferred,
				OccurredAt:     context.Input.EvaluatedAt,
				WalkID:         context.Input.WalkID,
				PlayerID:       context.Input.PlayerID,
				HexID:          &finalHexID,
				RuleSetVersion: context.RuleSetVersion,
				EngineVersion:  context.Input.EngineVersion,
			})
		}
	}

	if len(context.ScoreChanges) > 0 {
		events = append(events, DomainEvent{
			ID:             context.EventIDs[len(events)],
			Type:           EventTypePlayerScoreChanged,
			OccurredAt:     context.Input.EvaluatedAt,
			WalkID:         context.Input.WalkID,
			PlayerID:       context.Input.PlayerID,
			RuleSetVersion: context.RuleSetVersion,
			EngineVersion:  context.Input.EngineVersion,
		})
	}

	return events
}
