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

	// Restore Check
	//
	// 判断 source 是否为 registry-sync 生成的镜像
	//
	// 如果存在:
	//
	// org.registry-sync.source
	//
	// 则根据原始 source 恢复干净路径
	//
	if e.metadataResolver != nil {

		sourceMetadata, err :=
			e.metadataResolver.ResolveMetadata(
				ctx,
				source,
			)

		if err == nil &&
			sourceMetadata.Source != "" {

			restoreTarget :=
				BuildRestoreTarget(
					sourceMetadata.Source,
					target,
				)

			if restoreTarget != "" {

				targetImage = restoreTarget
			}
		}
	}

	// Metadata Check
	//
	// 判断目标镜像是否已经同步过相同 digest
	//
	if e.metadataResolver != nil {

		targetMetadata, err :=
			e.metadataResolver.ResolveMetadata(
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
