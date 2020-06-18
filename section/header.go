package section

import (
	"bufio"
	"fmt"
	"os"
)

type Header struct {
	data map[string]string
}

func (h *Header) init() {
	//Init data map if not done already
	if h.data == nil {
		h.data = map[string]string{
			"Full Name":   "",
			"Title":       "",
			"Address":     "",
			"Email":       "",
			"Web Address": "",
			"Twitter":     "",
		}
	}
}

//Edit allows a user edit this section
func (h *Header) Edit() error {
	h.init()
	fmt.Println("Editing an entire section will overwrite its data")
	for key := range h.data {
		fmt.Print(key + ": ")
		line, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			return err
		}
		h.data[key] = line
	}

	return nil
}
