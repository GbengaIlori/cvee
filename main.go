package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var (
	infoLogger  = log.New(os.Stdout, "INFO: ", log.LstdFlags)
	debugLogger = log.New(os.Stdout, "DEBUG: ", log.LstdFlags)
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
