package section

import (
	"bufio"
	"fmt"
	"os"
)

type Summary struct {
	summary string
}

var reader = bufio.NewReader(os.Stdin)

func (s *Summary) Edit() error {
	fmt.Print("Resume Objective / Summary: ")
	lines, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	s.summary = lines
	fmt.Println(s.summary)
	return nil
}
