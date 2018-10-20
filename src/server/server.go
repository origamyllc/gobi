package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"gobi/src/controllers"
)

func main() {
	// todo : add request timeouts
	// todo : separate out routes ?
	// todo:server needs a connection pool ?
	// todo secure the api endpoints
	router := mux.NewRouter()
	router.HandleFunc("/user", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/user/{id}", controllers.GetUser).Methods("GET")
	router.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE")
	router.HandleFunc("/user/{id}", controllers.UpdateUser).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", router))
}
