package main

import (
	"contact"
	"fmt"
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
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

	/*
		form, err := contact.NewForm("x", "x", "This is a message")
		// #fistname *lastname $content
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(form)
	*/
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	api.SetApp(rest.AppSimple(func(w rest.ResponseWriter, r *rest.Request) {
		w.WriteJson(map[string]string{"Body": "Hello World!"})
	}))
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}
