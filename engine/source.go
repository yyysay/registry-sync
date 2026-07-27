package engine

import (
	"registry-sync/model"
)

func ResolveSources(
	plan model.Plan,
) []string {

	return []string{
		BuildImageName(
			plan.Image.Registry,
			plan.Image.Repository,
			plan.Image.Tag,
		),
	}
}
