package main

import (
	"net/http"

	"github.com/jirawat050/api/controller"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
    r.GET("/hospital/:id", uc.GetUser)
    r.GET("/hospital", uc.GetAllUser)
	
	http.ListenAndServe("localhost:8080", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://api:paopao00@ds151814.mlab.com:51814/api")

	if err != nil {
		panic(err)
	}
	return s
}
