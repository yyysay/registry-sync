package engine

import (
	"context"
	"sync"

	"registry-sync/model"
)

func (e *Engine) ExecuteAll(
	ctx context.Context,
	plans []model.Plan,
	workers int,
) []model.ExecutionResult {

	sem := make(chan struct{}, workers)

	var wg sync.WaitGroup

	results := make(
		[]model.ExecutionResult,
		0,
		len(plans),
	)

	var mu sync.Mutex

	for _, plan := range plans {

		wg.Add(1)

		go func(plan model.Plan) {

			defer wg.Done()

			sem <- struct{}{}

			defer func() {
				<-sem
			}()

			err := e.Execute(
				ctx,
				plan,
			)

			result := model.ExecutionResult{

				Image: buildImageName(
					plan.Image,
				),

				Success: err == nil,

				Error: err,
			}

			mu.Lock()

			results = append(
				results,
				result,
			)

			mu.Unlock()

		}(plan)
	}

	wg.Wait()

	return results
}

func buildImageName(
	image model.Image,
) string {

	name := image.Repository

	if image.Tag != "" {

		name += ":" + image.Tag
	}

	return name
}
