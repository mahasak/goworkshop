package contact

import "gopkg.in/mgo.v2"

func ListAllMongo() []*Form {
	session, err := mgo.Dial("192.168.100.35:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	c := session.DB("contact").C("form")

	var forms []*Form
	c.Find(nil).All(&forms)

	return forms
}
