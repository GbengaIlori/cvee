package section

import (
	"fmt"
)

type Summary struct {
	summary string
}

func (s *Summary) Edit() error {
	fmt.Print("Resume Objective / Summary: ")
	lines, err := getLongString()
	if err != nil {
		return err
	}
	s.summary = lines
	return nil
}
