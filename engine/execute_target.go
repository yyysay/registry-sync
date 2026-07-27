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
	metadata model.ImageMetadata,
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

	targetImage := BuildTargetImage(
		plan.Image,
		target,
	)

	// Metadata Check
	//
	// 读取目标镜像 label
	//
	if e.metadataResolver != nil {

		targetMetadata, err := e.metadataResolver.ResolveMetadata(
			ctx,
			targetImage,
		)

		if err == nil {

			if targetMetadata.Digest == resolved.Digest {

				result.Success = true
				result.Cached = true

				return result
			}
		}
	}

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
		metadata,
	)

	if err != nil {

		result.Error = err

		return result
	}

	result.Success = true

	return result
}
