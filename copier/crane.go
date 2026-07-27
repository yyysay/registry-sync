package copier

import (
	"context"

	"registry-sync/engine"
	"registry-sync/model"

	"github.com/google/go-containerregistry/pkg/crane"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/mutate"
	"github.com/google/go-containerregistry/pkg/v1/remote"
)

type CraneCopier struct {
}

func New() *CraneCopier {

	return &CraneCopier{}
}

func (c *CraneCopier) Copy(
	ctx context.Context,
	source string,
	target string,
	platform []string,
	metadata model.ImageMetadata,
) error {

	opts := []crane.Option{
		crane.WithContext(ctx),
	}

	opts = append(
		opts,
		buildPlatformOption(platform)...,
	)

	image, err := crane.Pull(
		source,
		opts...,
	)

	if err != nil {

		dumpFailed(err)

		return err
	}

	configFile, err := image.ConfigFile()

	if err != nil {

		dumpFailed(err)

		return err
	}

	if configFile.Config.Labels == nil {

		configFile.Config.Labels = map[string]string{}
	}

	configFile.Config.Labels["org.registry-sync.source"] =
		metadata.Source

	configFile.Config.Labels["org.registry-sync.digest"] =
		metadata.Digest

	image, err = mutate.ConfigFile(
		image,
		configFile,
	)

	if err != nil {

		dumpFailed(err)

		return err
	}

	ref, err := name.ParseReference(
		target,
	)

	if err != nil {

		dumpFailed(err)

		return err
	}

	err = remote.Write(
		ref,
		image,
	)

	if err != nil {

		dumpFailed(err)

		return err
	}

	dumpSuccess()

	return nil
}

var _ engine.Copier = (*CraneCopier)(nil)
