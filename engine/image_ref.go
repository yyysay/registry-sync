package engine

import (
	"registry-sync/model"
)

func BuildImageRef(
	image model.Image,
) string {

	ref := image.Repository

	if image.Registry != "" {
		ref = image.Registry + "/" + ref
	}

	if image.Tag != "" {
		ref += ":" + image.Tag
	}

	return ref
}
