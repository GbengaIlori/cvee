package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "cvee",
		Usage: "Manage your resume from the terminal",
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
