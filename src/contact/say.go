package contact

// Form user defined type
type Form struct {
	Sender
	Content string
}

// Sender user defined type
type Sender struct {
	Firstname string
	Lastname  string
}

// toString format string
func (f Form) toString() string {
	return "#" + f.Firstname + "*" + f.Lastname + "$" + f.Content
}

// NewForm return form object
func NewForm(firstname string, lastname string, msg string) *Form {
	return &Form{
		Sender: Sender{
			Firstname: firstname,
			Lastname:  lastname,
		},
		Content: msg,
	}
}

// Say return Hi
func Say() string {
	return "Hi"
}
