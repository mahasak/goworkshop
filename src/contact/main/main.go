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

	var arr [5]int
	arr[0] = 10
	arr[0] = 5
	fmt.Println(arr)
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	router, _ := rest.MakeRouter(
		rest.Get("/form", listForms),
		rest.Get("/form/:id", listForms),
		rest.Post("/form", listForms),
	)
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

func listForms(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(contact.ListAll())
}

/*
func getForms(w rest.ResponseWriter, r *rest.Request) {
	id, _ := strconv.Atoi(r.PathParam("id"))
	w.WriteJson(forms[id])
}
func createForms(w rest.ResponseWriter, r *rest.Request) {
	var form contact.Form
	r.DecodeJsonPayload(&form)
	forms = append(forms, &form)
}
*/
