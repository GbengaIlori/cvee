package section

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Experience struct {
	data   map[string]string
	points []string
}

func (e *Experience) init() {
	//Init data map if not done already
	if e.data == nil {
		e.data = map[string]string{
			"Job Title":         "",
			"Work Place":        "",
			"Years (xxxx-xxxx)": "",
		}
	}
}

//Edit allows a user edit this section
func (e *Experience) Edit() error {
	e.init()
	for key := range e.data {
		fmt.Print(key + ": ")
		line, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			return err
		}
		e.data[key] = line
	}

	// populate points
	if err := e.editPoints(); err != nil {
		return err
	}
	return nil
}

func (e *Experience) editPoints() error {
	for {
		fmt.Print("Notable events / What you did (Empty string to stop): ")
		line, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			return err
		}
		if strings.TrimSpace(line) == "" {
			break
		}
		e.points = append(e.points, line)
	}

	fmt.Println(e.points)
	return nil
}
