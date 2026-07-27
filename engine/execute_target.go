package engine

import (
	"context"

	"registry-sync/model"
)

func (e *Engine) executeTarget(
	ctx context.Context,
	plan model.Plan,
	resolved model.ResolvedImage,
	source string,
	target model.Target,
) model.ExecutionResult {

	key := BuildCacheKey(
		resolved,
		target,
	)

	result := model.ExecutionResult{
		Image:  key,
		Target: target.Name,
	}

	// Cache Check
	if e.cache != nil {

		if e.cache.Check(
			ctx,
			key,
		) {

			result.Success = true
			result.Cached = true

			return result
		}
	}

	targetImage := BuildTargetImage(
		plan.Image,
		target,
	)

	dumpCopyTask(
		source,
		targetImage,
		plan.Image.Platform,
	)

	err := e.copier.Copy(
		ctx,
		source,
		targetImage,
		plan.Image.Platform,
	)

	if err != nil {

		result.Error = err

		return result
	}

	// Cache Save
	if e.cache != nil {

		err := e.cache.Save(
			ctx,
			key,
		)

		if err != nil {

			result.Error = err

			return result
		}
	}

	result.Success = true

	return result
}
