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
	form, err := contact.NewForm("Pondd", "NoobMe", "This is a message")
	// #fistname *lastname $content
	if err != nil {
		fmt.Println(form)
	}
	fmt.Println(form)
}
