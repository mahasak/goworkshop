package main

import (
	"contact"
	"fmt"
	"log"
	"net/http"
	"strconv"

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

	var arr [5]int
	arr[0] = 10
	arr[0] = 5
	fmt.Println(arr)
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	forms := []*contact.Form{
		&contact.Form{
			Sender: contact.Sender{
				Firstname: "Pondd1",
				Lastname:  "NoobMe",
			},
			Content: "Contentt",
		},
		&contact.Form{
			Sender: contact.Sender{
				Firstname: "Pondd2",
				Lastname:  "NoobMe",
			},
			Content: "Contentt",
		},
	}
	router, _ := rest.MakeRouter(
		rest.Get("/form", func(w rest.ResponseWriter, r *rest.Request) {
			w.WriteJson(forms)
		}),
		rest.Get("/form/:id", func(w rest.ResponseWriter, r *rest.Request) {
			id, _ := strconv.Atoi(r.PathParam("id"))
			w.WriteJson(forms[id])
		}),
		rest.Post("/form", func(w rest.ResponseWriter, r *rest.Request) {
			var form contact.Form
			r.DecodeJsonPayload(&form)
			forms = append(forms, &form)
		}),
	)
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}
