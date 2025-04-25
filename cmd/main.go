package main

import (
	"backend/cmd/app"
	"os"
)

func main() {
	command := app.NewCommand()
	err := command.Execute()
	if err != nil {
		os.Exit(1)
	}
}
