package section

type Header struct {
	Name       string
	Title      string
	Address    string
	Email      string
	WebAddress string
	Twitter    string
}

//Edit allows a user edit this section
func (h *Header) Edit() error {
}
