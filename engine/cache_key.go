package engine

import (
	"fmt"
	"strings"

	"registry-sync/model"
)

func BuildCacheKey(
	image model.ResolvedImage,
	target model.Target,
) string {

	return fmt.Sprintf(
		"%s/%s@%s:%s:%s",
		image.Registry,
		image.Repository,
		image.Digest,
		strings.Join(
			image.Platform,
			",",
		),
		target.Name,
	)
}
