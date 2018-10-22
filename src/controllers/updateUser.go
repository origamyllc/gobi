package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	queries "gobi/src/constants/sql"
	statics "gobi/src/constants/app"
	user_dao "gobi/src/dao/User"
	db "gobi/src/database/postgres"
	"log"
	"net/http"
	"os"
	"strconv"
)


func UpdateUser(w http.ResponseWriter, r *http.Request){
	// should check headers for the token and proceed only if token matches
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var data user_dao.User
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}

	params := mux.Vars(r)
	key:= params["id"]
	log.Println("Url Param 'key' is: " + string(key))
	i, err := strconv.Atoi(string(key))
	log.Println(data)
	conn, err := db.Connect();
	if err != nil {
		panic(err)
	}
	host, _ := os.Hostname();
	port := statics.PORT;
	query := fmt.Sprintf(queries.UPDATE_USER_NAME__QUERY, data.FirstName, data.LastName, i )

	err = db.Update(query,conn);
	if err == nil {
		log.Print("Host: "+host +"Port: " +port+ " result:= Updated")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"code":"200","message": "record sucessfully Updated"})
	} else {
		log.Print("Host: " + host + "Port: " + port + " result:= Not updated")
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(500) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	defer conn.Close();
}
