package main

import (
	"fmt"
	"strconv"
)

func ParseWorkers(
	args []string,
	defaultWorkers int,
) (int, error) {

	workers := defaultWorkers

	for i := 0; i < len(args); i++ {

		if args[i] != "--workers" {
			continue
		}

		if i+1 >= len(args) {

			return 0, fmt.Errorf(
				"--workers requires value",
			)
		}

		n, err := strconv.Atoi(
			args[i+1],
		)

		if err != nil {

			return 0, fmt.Errorf(
				"invalid workers value: %s",
				args[i+1],
			)
		}

		workers = n

		break
	}

	if workers <= 0 {

		return 0, fmt.Errorf(
			"workers must be greater than 0",
		)
	}

	return workers, nil
}
