// Functions here help to query the user
package section

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func populateSlice(message string, s *[]string) error {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(message)
		line, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		if strings.TrimSpace(line) == "" {
			break
		}
		*s = append(*s, line)
	}
	return nil
}

func populateMap(m *map[string]string) error {
	reader := bufio.NewReader(os.Stdin)
	for key := range *m {
		fmt.Print(key + ": ")
		line, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		(*m)[key] = line
	}
	return nil
}

func getLongString() (string, error) {
	lines, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return "", err
	}
	return lines, nil
}
