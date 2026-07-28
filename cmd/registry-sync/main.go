package main

import (
	"log"
	"os"
)

func main() {

	err := Execute(
		os.Args[1:],
	)

	if err != nil {
		log.Fatal(err)
	}
}
