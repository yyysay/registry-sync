package engine

import (
	"strings"

	"registry-sync/model"
)

// BuildRestoreTarget
//
// 根据 registry-sync metadata 中记录的原始 source
// 生成恢复后的干净目标地址.
//
// source:
//
//	ghcr.io/finb/bark-server:v2
//
// target:
//
//	registry.local.y3-3am.top
//
// result:
//
//	registry.local.y3-3am.top/finb/bark-server:v2
func BuildRestoreTarget(
	source string,
	target model.Target,
) string {

	repository, tag := splitRestoreSource(
		source,
	)

	if repository == "" {

		return ""
	}

	return BuildImageName(
		target.Registry,
		applyNamespace(
			target.Namespace,
			repository,
		),
		tag,
	)
}

func splitRestoreSource(
	source string,
) (
	string,
	string,
) {

	parts := strings.Split(
		source,
		"/",
	)

	if len(parts) < 2 {

		return "", ""
	}

	repository := strings.Join(
		parts[1:],
		"/",
	)

	tag := ""

	if index := strings.LastIndex(
		repository,
		":",
	); index > 0 {

		tag = repository[index+1:]

		repository = repository[:index]
	}

	return repository, tag
}

func applyNamespace(
	namespace string,
	repository string,
) string {

	if namespace == "" {

		return repository
	}

	return namespace + "/" + repository
}
