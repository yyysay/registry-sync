package engine

import (
	"context"

	"registry-sync/model"
)

type Copier interface {
	Copy(
		ctx context.Context,
		source string,
		target string,
		platform []string,
		metadata model.ImageMetadata,
	) error
}
