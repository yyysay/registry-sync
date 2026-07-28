package engine

import (
	"fmt"
	"strings"
)

func dumpCopyTask(
	source string,
	target string,
	platform []string,
) {

	task := source +
		" ==> " +
		target

	if len(platform) > 0 {

		task += " (" +
			strings.Join(
				platform,
				",",
			) +
			")"
	}

	fmt.Println(
		task,
	)
}
