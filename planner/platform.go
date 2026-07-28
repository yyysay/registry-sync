package planner

import "registry-sync/config"

func resolvePlatform(
	cfg *config.Config,
	source config.SourceConfig,
	image *config.ImageConfig,
) []string {

	// image 优先
	if image != nil && len(image.Platform) > 0 {
		return image.Platform
	}

	// source 第二
	if len(source.Platform) > 0 {
		return source.Platform
	}

	// global 最后
	return cfg.Platform
}
