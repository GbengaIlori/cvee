package section

import (
	"bufio"
	"fmt"
	"os"
)

type Header struct {
	Name       string
	Title      string
	Address    string
	Email      string
	WebAddress string
	Twitter    string
}

var reader = bufio.NewReader(os.Stdin)

//Edit allows a user edit this section
func (h *Header) Edit() error {
	// get name
	fmt.Print("Your full name: ")
	fname, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	h.Name = fname

	// get title
	fmt.Print("Your title, e.g. Software Engineer: ")
	title, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	h.Title = title

	fmt.Println(h)
	return nil
}
