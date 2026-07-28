package copier

import (
	"context"

	"registry-sync/model"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/crane"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/mutate"
)

func pullImage(
	ctx context.Context,
	source string,
	platform string,
	metadata model.ImageMetadata,
) (v1.Image, error) {

	opts := []crane.Option{
		crane.WithContext(ctx),

		crane.WithAuthFromKeychain(
			authn.DefaultKeychain,
		),
	}

	opts = append(
		opts,
		buildPlatformOption(
			platform,
		)...,
	)

	image, err :=
		crane.Pull(
			source,
			opts...,
		)

	if err != nil {

		return nil, err
	}

	// 获取当前实际 manifest digest
	//
	// 单平台:
	//
	// source index
	//      |
	//      v
	// linux/amd64 manifest digest
	//
	// 这里获取的是最终复制对象 digest
	imageDigest, err :=
		image.Digest()

	if err != nil {

		return nil, err
	}

	configFile, err :=
		image.ConfigFile()

	if err != nil {

		return nil, err
	}

	if configFile.Config.Labels == nil {

		configFile.Config.Labels =
			map[string]string{}
	}

	configFile.Config.Labels["org.registry-sync.version"] =
		"1"

	configFile.Config.Labels["org.registry-sync.source"] =
		metadata.Source

	configFile.Config.Labels["org.registry-sync.digest"] =
		imageDigest.String()

	image, err =
		mutate.ConfigFile(
			image,
			configFile,
		)

	if err != nil {

		return nil, err
	}

	return image, nil
}
