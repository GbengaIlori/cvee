package section

import (
	"fmt"
)

type Skills struct {
	data []string
}

func (s *Skills) Edit() error {
	fmt.Println(".................SKILLS.................")
	fmt.Println()

	err := populateSlice("Technologies & Skills (Empty string to stop): ", &s.data)
	if err != nil {
		return err
	}
	return nil
}
