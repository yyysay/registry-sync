package copier

import (
	"fmt"
)

func dumpSuccess() {

	fmt.Println(
		"SUCCESS",
	)

	fmt.Println()
}

func dumpFailed(
	err error,
) {

	fmt.Println(
		"FAILED",
	)

	fmt.Println(
		"ERROR:",
		err,
	)

	fmt.Println()
}
