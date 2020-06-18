package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func create() func(c *cli.Context) error {
	return func(c *cli.Context) error {
		fmt.Println("Create a new resume!")
		return nil
	}
}
