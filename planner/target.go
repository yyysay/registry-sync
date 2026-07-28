package planner

import (
	"registry-sync/config"
	"registry-sync/mapper"
	"registry-sync/model"
)

// resolveTargets 根据 sync 配置生成目标仓库列表
//
// 优先级：
// source.Sync > global Sync
func resolveTargets(
	cfg *config.Config,
	source config.SourceConfig,
) []model.Target {

	var names []string

	// source 级覆盖全局
	if len(source.Sync) > 0 {
		names = source.Sync
	} else {
		names = cfg.Sync
	}

	var targets []model.Target

	for _, name := range names {

		dest, ok := cfg.Dest[name]
		if !ok {
			// 暂时忽略不存在的目标
			// 后续 Validate 阶段处理
			continue
		}

		targets = append(targets, model.Target{
			Name:      name,
			Registry:  dest.Registry,
			Namespace: dest.Namespace,
			Flatten:   dest.Flatten,
			Auth:      mapper.ConvertAuth(dest.Auth),
		})
	}

	return targets
}
