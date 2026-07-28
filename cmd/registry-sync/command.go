package main

import "fmt"

func Execute(
	args []string,
) error {

	if len(args) == 0 {

		return fmt.Errorf(
			"usage: registry-sync [plan|sync] <config.yaml>",
		)
	}

	switch args[0] {

	case "plan":

		return PlanCommand(
			args[1:],
		)

	case "sync":

		return SyncCommand(
			args[1:],
		)

	default:

		// 兼容旧模式:
		// registry-sync config.yaml
		return SyncCommand(
			args,
		)
	}
}
