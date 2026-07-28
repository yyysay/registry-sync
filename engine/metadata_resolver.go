package engine

import (
	"context"
	"fmt"
	"strings"

	"registry-sync/model"

	"github.com/google/go-containerregistry/pkg/crane"
)

const RegistrySyncMetadataVersion = "1"

type MetadataResolver interface {
	ResolveMetadata(
		ctx context.Context,
		image string,
	) (model.ImageMetadata, error)
}

type CraneMetadataResolver struct {
}

func NewCraneMetadataResolver() *CraneMetadataResolver {

	return &CraneMetadataResolver{}
}

func (r *CraneMetadataResolver) ResolveMetadata(
	ctx context.Context,
	image string,
) (model.ImageMetadata, error) {

	img, err := crane.Pull(
		image,
		crane.WithContext(ctx),
	)

	if err != nil {

		return model.ImageMetadata{}, err
	}

	configFile, err := img.ConfigFile()

	if err != nil {

		return model.ImageMetadata{}, err
	}

	labels := configFile.Config.Labels

	if labels == nil {

		return model.ImageMetadata{}, fmt.Errorf(
			"image has no registry-sync metadata",
		)
	}

	version := labels["org.registry-sync.version"]

	if version != RegistrySyncMetadataVersion {

		return model.ImageMetadata{}, fmt.Errorf(
			"unsupported registry-sync metadata version: %s",
			version,
		)
	}

	platform := []string{}

	if value := labels["org.registry-sync.platform"]; value != "" {

		for _, item := range strings.Split(
			value,
			",",
		) {

			platform = append(
				platform,
				strings.TrimSpace(item),
			)
		}
	}

	return model.ImageMetadata{

		Source: labels["org.registry-sync.source"],

		Digest: labels["org.registry-sync.digest"],

		Platform: platform,
	}, nil
}

var _ MetadataResolver = (*CraneMetadataResolver)(nil)
