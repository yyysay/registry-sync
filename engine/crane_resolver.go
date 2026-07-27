package engine

import (
	"context"

	"registry-sync/model"

	"github.com/google/go-containerregistry/pkg/crane"
)

type CraneResolver struct {
}

func NewCraneResolver() *CraneResolver {

	return &CraneResolver{}
}

func (r *CraneResolver) Resolve(
	ctx context.Context,
	source string,
	image model.Image,
) (model.ResolvedImage, error) {

	digest, err := crane.Digest(
		source,
		crane.WithContext(ctx),
	)

	if err != nil {

		return model.ResolvedImage{}, err
	}

	return model.ResolvedImage{

		Registry: image.Registry,

		Repository: image.Repository,

		Tag: image.Tag,

		Digest: digest,

		Platform: image.Platform,
	}, nil
}

var _ Resolver = (*CraneResolver)(nil)
