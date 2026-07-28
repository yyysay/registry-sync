package engine

import (
	"context"

	"registry-sync/model"
)

func (e *Engine) Execute(
	ctx context.Context,
	plan model.Plan,
) []model.ExecutionResult {

	resolved, source, err :=
		e.resolveSource(
			ctx,
			plan,
		)

	if err != nil {

		return []model.ExecutionResult{
			{
				Image: BuildImageRef(plan.Image),
				Error: err,
			},
		}
	}

	if source == "" {

		return []model.ExecutionResult{
			{
				Image: BuildImageRef(plan.Image),
			},
		}
	}

	results := make(
		[]model.ExecutionResult,
		0,
		len(plan.Targets),
	)

	metadata := model.ImageMetadata{
		Source:   source,
		Digest:   resolved.Digest,
		Platform: resolved.Platform,
	}

	for _, target := range plan.Targets {

		result :=
			e.executeTarget(
				ctx,
				plan,
				resolved,
				source,
				metadata,
				target,
			)

		results = append(
			results,
			result,
		)
	}

	return results
}
