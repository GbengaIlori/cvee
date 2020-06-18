package section

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Skills struct {
	data []string
}

func (s *Skills) Edit() error {
	for {
		fmt.Print("Technologies and Skills (Empty string to stop): ")
		line, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			return err
		}
		if strings.TrimSpace(line) == "" {
			break
		}
		s.data = append(s.data, line)
	}

	fmt.Println(s.data)
	return nil
}
