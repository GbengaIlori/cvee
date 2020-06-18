package section

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
	if err := populateMap(&e.data); err != nil {
		return err
	}

	// populate points
	err := populateSlice("Notable events / What you did (Empty string to stop): ", &e.points)

	if err != nil {
		return err
	}
	return nil
}
