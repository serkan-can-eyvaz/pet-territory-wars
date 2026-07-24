package gameengine

// resolveScoreChanges converts a territory change into its canonical score deltas.
func resolveScoreChanges(input ResolveWalkInput, change HexChange) []ScoreChange {
	scoreChanges := make([]ScoreChange, 0)

	switch change.ChangeType {
	case HexChangeTypeEmptyCapture:
		scoreChanges = append(scoreChanges,
			ScoreChange{
				PlayerID: input.PlayerID,
				Metric:   ScoreMetricActiveHex,
				Delta:    1,
				Reason:   ScoreChangeReasonEmptyCapture,
				WalkID:   input.WalkID,
			},
			ScoreChange{
				PlayerID: input.PlayerID,
				Metric:   ScoreMetricCapture,
				Delta:    1,
				Reason:   ScoreChangeReasonEmptyCapture,
				WalkID:   input.WalkID,
			},
			ScoreChange{
				PlayerID: input.PlayerID,
				Metric:   ScoreMetricLifetimeCapture,
				Delta:    1,
				Reason:   ScoreChangeReasonEmptyCapture,
				WalkID:   input.WalkID,
			},
		)
	case HexChangeTypeOwnerDefense:
		scoreChanges = append(scoreChanges, ScoreChange{
			PlayerID: input.PlayerID,
			Metric:   ScoreMetricDefense,
			Delta:    1,
			Reason:   ScoreChangeReasonOwnerDefense,
			WalkID:   input.WalkID,
		})
	case HexChangeTypeOwnershipTransfer:
		scoreChanges = append(scoreChanges,
			ScoreChange{
				PlayerID: input.PlayerID,
				Metric:   ScoreMetricActiveHex,
				Delta:    1,
				Reason:   ScoreChangeReasonOwnershipTransfer,
				WalkID:   input.WalkID,
			},
			ScoreChange{
				PlayerID: input.PlayerID,
				Metric:   ScoreMetricSteal,
				Delta:    1,
				Reason:   ScoreChangeReasonOwnershipTransfer,
				WalkID:   input.WalkID,
			},
			ScoreChange{
				PlayerID: input.PlayerID,
				Metric:   ScoreMetricLifetimeSteal,
				Delta:    1,
				Reason:   ScoreChangeReasonOwnershipTransfer,
				WalkID:   input.WalkID,
			},
			ScoreChange{
				PlayerID: *change.PreviousOwnerID,
				Metric:   ScoreMetricActiveHex,
				Delta:    -1,
				Reason:   ScoreChangeReasonOwnershipTransfer,
				WalkID:   input.WalkID,
			},
		)
	}

	return scoreChanges
}
