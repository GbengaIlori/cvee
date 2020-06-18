package section

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
	if err := populateMap(&h.data); err != nil {
		return err
	}
	return nil
}
