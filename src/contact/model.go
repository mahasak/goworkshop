package contact

import (
	"errors"
	"fmt"
)

var forms = []*Form{
	&Form{
		Sender: Sender{
			Firstname: "Pondd1",
			Lastname:  "NoobMe",
		},
		Content: "Contentt",
	},
	&Form{
		Sender: Sender{
			Firstname: "Pondd2",
			Lastname:  "NoobMe",
		},
		Content: "Contentt",
	},
}

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

// String for sender
func (s Sender) String() string {
	return "#" + s.Firstname + "*" + s.Lastname
}

// String format string
func (f Form) String() string {
	return fmt.Sprintf("%s + %s", f.Sender, f.Content)
	//return " #" + f.Firstname + " *" + f.Lastname + " $" + f.Content
}

// NewForm return form object
func NewForm(firstname, lastname, msg string) (*Form, error) {
	if len(firstname) < 2 && len(lastname) < 2 {
		return nil, errors.New("สั้นไปทุกอย่าง")
	}
	return &Form{
		Sender: Sender{
			Firstname: firstname,
			Lastname:  lastname,
		},
		Content: msg,
	}, nil
}

//ListAll return forms objects
func ListAll() []*Form {
	return forms
}
