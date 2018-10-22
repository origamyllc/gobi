package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	statics "gobi/src/constants/app"
	queries "gobi/src/constants/sql"
	db "gobi/src/database/postgres"
	JwtToken "gobi/src/dao/Login"
	"log"
	"net/http"
	"os"
)

func DeleteUser(w http.ResponseWriter, r *http.Request){
	// should check headers for the token and proceed only if token matches
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	xauth := r.Header.Get("x-authorizarion")
	params := mux.Vars(r)
	key:= params["id"]
	log.Println("Url Param 'key' is: " + string(key))
	conn, err := db.Connect();
	if err != nil {
		panic(err)
	}
	host, _ := os.Hostname();
	port := statics.PORT;
	sqlStatement := queries.DELETE_USER_BY_ID_QUERY + string(key);
	tokenquery := fmt.Sprintf(queries.GET_USER_TOKEN_QUERY,xauth);
	var tokenDao JwtToken.JwtToken;
	var tokens  []JwtToken.JwtToken;
	rows,err := db.Get(tokenquery,conn)
	for rows.Next() {
		error :=  rows.Scan(&tokenDao.Token)
		if error != nil {
			log.Fatal("Host: "+host +"Port: " +port)
			log.Fatal(error)
			panic(error)
		} else{
			tokens = append(tokens,tokenDao);
		}
	}
	m,err := json.Marshal(tokens);
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(401) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	log.Print("Host: "+host +"Port: " +port+ " result:= "+string(m))
	if string(m) != "null" {
		error := db.Delete(sqlStatement, conn)
		fmt.Print(error)
		if error == nil {
			log.Print("Host: " + host + "Port: " + port + " result:= Deleted")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(map[string]string{"code": "200", "message": "record sucessfully deleted"})
		} else {
			log.Fatal("Host: " + host + "Port: " + port + " result:= Not Deleted")
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(500) // unprocessable entity
			if err := json.NewEncoder(w).Encode(err); err != nil {
				panic(err)
			}
		}
	}else {

		w.WriteHeader(401) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	defer rows.Close();
	defer conn.Close();
}

