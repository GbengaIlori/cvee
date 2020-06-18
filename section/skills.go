package section

import (
	"fmt"
)

type Skills struct {
	data []string
}

func (s *Skills) Edit() error {
	err := populateSlice("Technologies & Skills (Empty string to stop): ", &s.data)
	if err != nil {
		return err
	}
	fmt.Println(s.data)
	return nil
}
