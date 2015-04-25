package main

import (
	"contact"
	"fmt"
)

func main() {
	msg := contact.Say()
	fmt.Println(msg)

	/*
		contact := contact.Form{
			Sender: contact.Sender{
				Firstname: "Pondd",
				Lastname:  "NoobMe",
			},
			Content: "this is a message",
		}
	*/
	form := contact.NewForm("Pondd", "NoobMe", "This is a message")
	fmt.Println(form)
}
