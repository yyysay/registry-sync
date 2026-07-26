package engine

import (
	"context"

	"registry-sync/model"
)

func (e *Engine) Execute(
	ctx context.Context,
	plan model.Plan,
) model.ExecutionResult {

	key := BuildCacheKey(
		plan,
	)

	result := model.ExecutionResult{
		Image: key,
	}

	// Cache Check
	//
	// hit:
	//     skip copy
	//
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

	sources := ResolveSources(plan)

	for _, target := range plan.Targets {

		targetImage := BuildTargetImage(
			plan.Image,
			target,
		)

		var lastErr error

		for _, source := range sources {

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

			if err == nil {

				lastErr = nil

				break
			}

			lastErr = err
		}

		if lastErr != nil {

			result.Success = false
			result.Error = lastErr

			return result
		}
	}

	// Cache Save
	//
	// copy 全部成功后记录
	//
	if e.cache != nil {

		err := e.cache.Save(
			ctx,
			key,
		)

		if err != nil {

			result.Success = false
			result.Error = err

			return result
		}
	}

	result.Success = true

	return result
}
