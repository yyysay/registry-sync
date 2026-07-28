package copier

import (
	"context"
	"strings"

	"registry-sync/model"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/mutate"
	"github.com/google/go-containerregistry/pkg/v1/remote"
)

func (c *CraneCopier) copyMulti(
	ctx context.Context,
	source string,
	target string,
	platforms []string,
	metadata model.ImageMetadata,
) error {

	var manifests []mutate.IndexAddendum

	for _, platform := range platforms {

		image, err :=
			pullImage(
				ctx,
				source,
				platform,
				metadata,
			)

		if err != nil {

			dumpFailed(err)

			return err
		}

		parts :=
			strings.Split(
				platform,
				"/",
			)

		var platformInfo *v1.Platform

		if len(parts) == 2 {

			platformInfo =
				&v1.Platform{
					OS:           parts[0],
					Architecture: parts[1],
				}
		}

		descriptor := v1.Descriptor{
			Platform: platformInfo,
		}

		manifests = append(
			manifests,
			mutate.IndexAddendum{
				Add: image,

				Descriptor: descriptor,
			},
		)
	}

	index :=
		mutate.AppendManifests(
			emptyIndex(),
			manifests...,
		)

	ref, err :=
		name.ParseReference(
			target,
		)

	if err != nil {

		dumpFailed(err)

		return err
	}

	err =
		remote.WriteIndex(
			ref,
			index,
			remote.WithContext(ctx),
			remote.WithAuthFromKeychain(
				authn.DefaultKeychain,
			),
		)

	if err != nil {

		dumpFailed(err)

		return err
	}

	dumpSuccess()

	return nil
}
