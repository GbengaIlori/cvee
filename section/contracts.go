package section

//Section unifies all resume sections
type Section interface {
	Edit() error
}
