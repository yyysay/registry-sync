package main

import (
	"fmt"

	"registry-sync/config"
	"registry-sync/planner"
)

func PlanCommand(
	args []string,
) error {

	if len(args) < 1 {

		return fmt.Errorf(
			"usage: registry-sync plan <config.yaml>",
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

	return nil
}
