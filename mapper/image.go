package mapper

import "registry-sync/model"

func ConvertImage(
	registry string,
	repository string,
	tag string,
	platform []string,
) model.Image {

	return model.Image{
		Registry:   registry,
		Repository: repository,
		Tag:        tag,
		Platform:   platform,
	}
}
