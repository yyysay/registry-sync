package engine

import (
	"context"

	"registry-sync/model"
)

type Resolver interface {
	Resolve(
		ctx context.Context,
		image model.Image,
	) (model.ResolvedImage, error)
}
