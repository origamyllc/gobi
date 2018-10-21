package main

import (
	"github.com/gorilla/mux"
	"gobi/src/controllers"
	"log"
	"net/http"
	"os"
)


func main() {

	// todo:server needs a connection pool ?
	// todo secure the api endpoints

	router := mux.NewRouter()
	os.Setenv("PORT", ":8000")
	router.HandleFunc("/user", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/user/{id}", controllers.GetUser).Methods("GET")
	router.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE")
	router.HandleFunc("/user/{id}", controllers.UpdateUser).Methods("PUT")
	log.Fatal(http.ListenAndServe( os.Getenv("PORT"), router))
}
