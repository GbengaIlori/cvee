package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	// Register subcommands and flags
	app := &cli.App{
		Name:  "cvee",
		Usage: "Manage your resume from the terminal",
		Commands: []*cli.Command{
			{
				Name:   "create",
				Usage:  "Start a new resume",
				Action: create(),
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
