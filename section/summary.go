package section

import (
	"bufio"
	"fmt"
	"os"
)

type Summary struct {
	summary string
}

func (s *Summary) Edit() error {
	fmt.Print("Resume Objective / Summary: ")
	lines, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return err
	}
	s.summary = lines
	fmt.Println(s.summary)
	return nil
}
