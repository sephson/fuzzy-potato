package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"

	"github.com/sephson/fuzzy-potato/controllers"
)
 


func main(){
	r := httprouter.New() //r is a variable
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser )
	r.POST("/user", uc.CreateUser )
	r.DELETE("/user/:id", uc.DeleteUser )

	http.ListenAndServe("localhost:8000", r)
}


func getSession() *mgo.Session{ //getsession function returns a pointer to mongodb session, this is used to dial the mongo url, below
	s, err := mgo.Dial("mongodb://localhost:27107")
	if err != nil{ //handling error
		panic(err)
	}
	return s

}