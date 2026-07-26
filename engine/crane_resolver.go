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
	image model.Image,
) (model.ResolvedImage, error) {

	ref := image.Registry + "/" + image.Repository

	if image.Tag != "" {

		ref += ":" + image.Tag
	}

	digest, err := crane.Digest(
		ref,
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
