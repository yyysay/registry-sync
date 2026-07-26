package main

import (
	"context"
	"fmt"

	"registry-sync/config"
	"registry-sync/copier"
	"registry-sync/engine"
	"registry-sync/planner"
)

func SyncCommand(
	args []string,
) error {

	workers, err := ParseWorkers(
		args,
		2,
	)

	if err != nil {
		return err
	}

	if len(args) < 1 {

		return fmt.Errorf(
			"usage: registry-sync sync <config.yaml>",
		)
	}

	cfg, err := config.LoadConfig(
		args[0],
	)

	if err != nil {
		return err
	}

	plans := planner.Build(cfg)

	fmt.Println(
		planner.Dump(plans),
	)

	e := engine.New(
		copier.New(),
		nil,
	)

	ctx := context.Background()

	results := e.ExecuteAll(
		ctx,
		plans,
		workers,
	)

	var failed int

	fmt.Println()
	fmt.Println("SYNC RESULT")

	for _, result := range results {

		if result.Cached {

			fmt.Println(
				"CACHED:",
				result.Image,
			)

			continue
		}

		if result.Success {

			fmt.Println(
				"SUCCESS:",
				result.Image,
			)

			continue
		}

		failed++

		fmt.Println(
			"FAILED:",
			result.Image,
		)

		if result.Error != nil {

			fmt.Println(
				"  ERROR:",
				result.Error,
			)
		}
	}

	if failed > 0 {

		return fmt.Errorf(
			"sync failed: %d task(s) failed",
			failed,
		)
	}

	return nil
}
