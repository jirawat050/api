package main

import (
	"net/http"

	"github.com/jirawat050/api/controller"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/hospital/:id", uc.GetUser)
    	r.GET("/hospital", uc.GetAllUser)
    	r.GET("/", uc.GetAllUser)

	http.ListenAndServe(":"+port, r)
}
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://api:paopao00@ds151814.mlab.com:51814/api")

	if err != nil {
		panic(err)
	}
	return s
}
