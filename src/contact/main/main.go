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
	w.WriteJson(contact.ListAllMongo())
}

func getForms(w rest.ResponseWriter, r *rest.Request) {
	id, _ := strconv.Atoi(r.PathParam("id"))
	w.WriteJson(contact.Get(id))
}

/*
func createForms(w rest.ResponseWriter, r *rest.Request) {
	var form contact.Form
	r.DecodeJsonPayload(&form)
	forms = append(forms, &form)
}
*/
