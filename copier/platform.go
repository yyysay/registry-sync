package copier

import (
	"strings"

	"github.com/google/go-containerregistry/pkg/crane"
	v1 "github.com/google/go-containerregistry/pkg/v1"
)

func buildPlatformOption(
	platform []string,
) []crane.Option {

	var opts []crane.Option

	// 多平台不指定 WithPlatform
	// 保留 manifest list
	if len(platform) != 1 {
		return opts
	}

	parts := strings.Split(
		platform[0],
		"/",
	)

	if len(parts) != 2 {
		return opts
	}

	opts = append(
		opts,
		crane.WithPlatform(
			&v1.Platform{
				OS:           parts[0],
				Architecture: parts[1],
			},
		),
	)

	return opts
}
