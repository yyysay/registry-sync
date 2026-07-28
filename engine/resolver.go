package engine

import (
	"context"

	"registry-sync/model"
)

type Resolver interface {
	Resolve(
		ctx context.Context,
		source string,
		image model.Image,
	) (model.ResolvedImage, error)
}
