package section

import (
	"fmt"
)

type Summary struct {
	summary string
}

func (s *Summary) Edit() error {
	fmt.Println(".................SUMMARY.................")
	fmt.Println()
	fmt.Print("Resume Objective / Summary: ")
	lines, err := getLongString()
	if err != nil {
		return err
	}
	s.summary = lines
	return nil
}
