package copier

import (
	"context"

	"registry-sync/engine"
	"registry-sync/model"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
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

	// 多平台:
	//
	// linux/amd64
	// linux/arm64
	//
	// 生成 OCI Image Index
	//
	if len(platform) > 1 {

		return c.copyMulti(
			ctx,
			source,
			target,
			platform,
			metadata,
		)
	}

	// 单平台:
	//
	// 指定 platform
	// 只复制对应架构

	if len(platform) == 1 {

		return c.copySingle(
			ctx,
			source,
			target,
			platform[0],
			metadata,
		)
	}

	// 未指定 platform
	//
	// 保持原行为:
	//
	// 保留完整 manifest list

	image, err :=
		pullImage(
			ctx,
			source,
			"",
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

	err =
		remote.Write(
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

var _ engine.Copier = (*CraneCopier)(nil)
