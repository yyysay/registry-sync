package planner

import (
	"registry-sync/config"
	"registry-sync/mapper"
	"registry-sync/model"
)

func Build(cfg *config.Config) []model.Plan {

	var plans []model.Plan

	for registry, source := range cfg.Sources {

		for repo, image := range source.Images {

			var tags []string

			platform := resolvePlatform(cfg, source, image)

			if image == nil {
				tags = []string{"latest"}
			} else {
				tags = image.Tags
			}

			if len(tags) == 0 {
				tags = []string{"latest"}
			}

			for _, tag := range tags {

				plans = append(plans, model.Plan{
					Image: mapper.ConvertImage(
						registry,
						repo,
						tag,
						platform,
					),

					Targets: resolveTargets(
						cfg,
						source,
					),
				})
			}
		}
	}

	return plans
}
