package controllers

import (
	"encoding/json"
	"fmt"
	queries "gobi/src/constants/sql"
	statics "gobi/src/constants/app"
	user_dao "gobi/src/dao/User"
	db "gobi/src/database/postgres"
	"log"
	"net/http"
	"os"
)

/**
Post Eg:
{"age":31,
 "email":"agghlhvvjjoun",
 "FirstName":"jhoffony",
 "LastName":"brave"}
 */
func CreateUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var data user_dao.User
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}


	log.Println(data)
	conn, err := db.Connect();
	if err != nil {
		panic(err)
	}
	host, _ := os.Hostname();
	port := statics.PORT;

	query := fmt.Sprintf(queries.ADD_USER_QUERY, data.Age, data.Email, data.FirstName, data.LastName )

	err = db.Insert(query,conn)
	if err == nil {
		log.Print("Host: "+host +"Port: " +port+ " result:= Inserted")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"code":"200","message": "record sucessfully inserted"})
	} else {
		log.Print("Host: " + host + "Port: " + port + " result:= Not inserted")
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(500) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	defer conn.Close();
}

