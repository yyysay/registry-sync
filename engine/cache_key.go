package engine

import (
	"fmt"
	"strings"

	"registry-sync/model"
)

func BuildCacheKey(
	plan model.Plan,
) string {

	return fmt.Sprintf(
		"%s:%s:%s",
		plan.Image.Repository,
		plan.Image.Tag,
		strings.Join(
			plan.Image.Platform,
			",",
		),
	)
}
