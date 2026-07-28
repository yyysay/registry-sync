package copier

import (
	"context"

	"registry-sync/model"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote"
)

func (c *CraneCopier) copySingle(
	ctx context.Context,
	source string,
	target string,
	platform string,
	metadata model.ImageMetadata,
) error {

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

	ref, err :=
		name.ParseReference(
			target,
		)

	if err != nil {

		dumpFailed(err)

		return err
	}

	err = remote.Write(
		ref,
		image,
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
