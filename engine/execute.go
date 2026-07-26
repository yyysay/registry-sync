package engine

import (
	"context"

	"registry-sync/model"
)

func (e *Engine) Execute(
	ctx context.Context,
	plan model.Plan,
) []model.ExecutionResult {

	results := make(
		[]model.ExecutionResult,
		0,
		len(plan.Targets),
	)

	resolved := model.ResolvedImage{
		Registry:   plan.Image.Registry,
		Repository: plan.Image.Repository,
		Tag:        plan.Image.Tag,
		Platform:   plan.Image.Platform,
	}

	if e.resolver != nil {

		image, err := e.resolver.Resolve(
			ctx,
			plan.Image,
		)

		if err != nil {

			return []model.ExecutionResult{
				{
					Image:   BuildImageRef(plan.Image),
					Success: false,
					Error:   err,
				},
			}
		}

		resolved = image
	}

	sources := ResolveSources(plan)

	for _, target := range plan.Targets {

		key := BuildCacheKey(
			resolved,
			target,
		)

		result := model.ExecutionResult{
			Image:  key,
			Target: target.Name,
		}

		// Cache Check
		//
		// hit:
		//     skip this target
		//
		if e.cache != nil {

			if e.cache.Check(
				ctx,
				key,
			) {

				result.Success = true
				result.Cached = true

				results = append(
					results,
					result,
				)

				continue
			}
		}

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

			results = append(
				results,
				result,
			)

			continue
		}

		// Cache Save
		//
		// 单个 target 成功后记录
		//
		if e.cache != nil {

			err := e.cache.Save(
				ctx,
				key,
			)

			if err != nil {

				result.Success = false
				result.Error = err

				results = append(
					results,
					result,
				)

				continue
			}
		}

		result.Success = true

		results = append(
			results,
			result,
		)
	}

	return results
}
