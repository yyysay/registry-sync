package engine

import (
	"context"
)

type Cache interface {
	Check(
		ctx context.Context,
		key string,
	) bool

	Save(
		ctx context.Context,
		key string,
	) error
}
