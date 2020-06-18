package section

import "fmt"

type Education struct {
	data       map[string]string
	coursework []string
}

func (e *Education) init() {
	//Init data map if not done already
	if e.data == nil {
		e.data = map[string]string{
			"Degree":            "",
			"School":            "",
			"Years (xxxx-xxxx)": "",
			"CGPA":              "",
		}
	}
}

//Edit allows a user edit this section
func (e *Education) Edit() error {
	fmt.Println(".................EDUCATION.................")
	fmt.Println()

	e.init()
	if err := populateMap(&e.data); err != nil {
		return err
	}

	// populate coursework
	err := populateSlice("Relevant Coursework (Empty string to stop): ", &e.coursework)
	if err != nil {
		return err
	}
	return nil
}
