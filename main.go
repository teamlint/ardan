package main

import (
	"log"

	"github.com/teamlint/ardan/cli/command"
)

func main() {
	err := command.Run()
	if err != nil {
		log.Fatal(err)
	}
}
