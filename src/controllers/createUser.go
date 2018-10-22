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
{
  "Username": "1username",
  "password": "1password",
  "secret": "secret",
  "FirstName": "first_name",
  "LastName": "last_name",
  "Email": "email",
  "OrganizationName": "organization_name"
}
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

	query := fmt.Sprintf(queries.ADD_USER_QUERY, data.Username,data.Password,data.Secret, data.FirstName, data.LastName, data.OrganizationName,data.Email )

	err= db.Insert(query,conn)
	fmt.Print(err)
	if err == nil {
		log.Print("Host: "+host +"Port: " +port+ " result:= Inserted")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"code":"200","message": "record sucessfully inserted"})
	} else {
		log.Print("Host: " + host + "Port: " + port + " result:= Not inserted")
		w.WriteHeader(500) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	defer conn.Close();
}

