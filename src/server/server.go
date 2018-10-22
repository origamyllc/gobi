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

	router := mux.NewRouter()
	os.Setenv("PORT", ":8000")
	router.HandleFunc("/user", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/user/{id}", controllers.GetUser).Methods("GET")
	router.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user/{id}", controllers.DeleteUser).Methods("DELETE")
	router.HandleFunc("/user/{id}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/token", controllers.CreateUserToken).Methods("POST")
	router.HandleFunc("/authenticate", controllers.GetUserToken).Methods("POST")
	log.Fatal(http.ListenAndServe( ":8000", router))
}
