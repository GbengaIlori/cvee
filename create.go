package main

import (
	"github.com/thealamu/cvee/section"
	"github.com/urfave/cli/v2"
)

func create() func(c *cli.Context) error {
	return func(c *cli.Context) error {
		// edit resume sections one after the other
		// this sections list is defined and maintained in main.go
		return editSections(sections)
	}
}

func editSections(sections []section.Section) error {
	for _, s := range sections {
		if err := s.Edit(); err != nil {
			return err
		}
	}
	return nil
}
