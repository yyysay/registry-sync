package engine

import (
	"context"

	"registry-sync/model"
)

func (e *Engine) resolveSource(
	ctx context.Context,
	plan model.Plan,
) (
	model.ResolvedImage,
	string,
	error,
) {

	resolved := model.ResolvedImage{
		Registry:   plan.Image.Registry,
		Repository: plan.Image.Repository,
		Tag:        plan.Image.Tag,
		Platform:   plan.Image.Platform,
	}

	sources := ResolveSources(plan)

	if len(sources) == 0 {

		return resolved, "", nil
	}

	if e.resolver == nil {

		return resolved, sources[0], nil
	}

	var lastErr error

	for _, source := range sources {

		image, err := e.resolver.Resolve(
			ctx,
			source,
			plan.Image,
		)

		if err == nil {

			return image, source, nil
		}

		lastErr = err
	}

	return resolved, "", lastErr
}
