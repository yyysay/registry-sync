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
) error {

	sem := make(chan struct{}, workers)

	var wg sync.WaitGroup

	var firstErr error
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

			if err != nil {

				mu.Lock()

				if firstErr == nil {
					firstErr = err
				}

				mu.Unlock()
			}

		}(plan)
	}

	wg.Wait()

	return firstErr
}
