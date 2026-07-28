package engine

import (
	"context"
	"strings"

	"registry-sync/model"

	"github.com/google/go-containerregistry/pkg/crane"
	v1 "github.com/google/go-containerregistry/pkg/v1"
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

	opts := []crane.Option{
		crane.WithContext(ctx),
	}

	// 如果指定单个平台:
	//
	// 例如:
	//
	// linux/amd64
	//
	// 则拉取对应 manifest
	// 获取真实 platform digest
	if len(image.Platform) == 1 {

		parts := strings.Split(
			image.Platform[0],
			"/",
		)

		if len(parts) == 2 {

			opts = append(
				opts,
				crane.WithPlatform(
					&v1.Platform{
						OS:           parts[0],
						Architecture: parts[1],
					},
				),
			)
		}
	}

	img, err := crane.Pull(
		source,
		opts...,
	)

	if err != nil {

		return model.ResolvedImage{}, err
	}

	digest, err :=
		img.Digest()

	if err != nil {

		return model.ResolvedImage{}, err
	}

	return model.ResolvedImage{

		Source: source,

		Registry: image.Registry,

		Repository: image.Repository,

		Tag: image.Tag,

		Digest: digest.String(),

		Platform: image.Platform,
	}, nil
}

var _ Resolver = (*CraneResolver)(nil)
