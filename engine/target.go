package engine

import (
	"registry-sync/model"
)

func BuildTargetImage(
	image model.Image,
	target model.Target,
) string {

	repository := image.Repository

	if target.Flatten {
		repository = flattenRepository(repository)
	}

	if target.Namespace != "" {
		repository = target.Namespace + "/" + repository
	}

	return BuildImageName(
		target.Registry,
		repository,
		image.Tag,
	)
}
