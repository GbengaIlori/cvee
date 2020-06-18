package main

import (
	"log"
	"os"

	sect "github.com/thealamu/cvee/section"
	"github.com/urfave/cli/v2"
)

// Create a list of all the sections
//TODO: Fill up with other sections
var sections = []sect.Section{
	&sect.Header{},
	&sect.Summary{},
	&sect.Skills{},
	&sect.Experience{},
}

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
