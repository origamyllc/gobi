package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	statics "gobi/src/constants/app"
	queries "gobi/src/constants/sql"
	user_dao "gobi/src/dao/User"
	db "gobi/src/database/postgres"
	JwtToken "gobi/src/dao/Login"
	"log"
	"net/http"
	"os"
)
func GetUser(w http.ResponseWriter, r *http.Request) {
	// should check headers for the token and proceed only if token matches
	w.Header().Set("Content-Type", "application/json")
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
	sqlStatement := queries.GET_USER_BY_ID_QUERY + string(key);
	var userDao user_dao.UserResponse;
	var users []user_dao.UserResponse;
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
		rows, err := db.Get(sqlStatement, conn)
		for rows.Next() {
			error := rows.Scan(&userDao.ID, &userDao.Username, &userDao.Password, &userDao.Secret, &userDao.FirstName, &userDao.LastName, &userDao.Email, &userDao.OrganizationName, &userDao.Created, &userDao.Updated)
			if error != nil {
				log.Fatal("Host: " + host + "Port: " + port)
				log.Fatal(error)
				panic(error)
			} else {
				users = append(users, userDao);
			}
		}
		m, err := json.Marshal(users);
		if err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(422) // unprocessable entity
			if err := json.NewEncoder(w).Encode(err); err != nil {
				panic(err)
			}
		}
		log.Print("Host: " + host + "Port: " + port + " result:= " + string(m))
		if string(m) != "null" {
			log.Print("Host: " + host + "Porsdt: " + port + " result:= " + string(m))
			json.NewEncoder(w).Encode(map[string]string{"code": "200", "result": string(m)})
		} else {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(500) // unprocessable entity
			if err := json.NewEncoder(w).Encode(err); err != nil {
				panic(err)
			}
		}
		defer rows.Close()
	}else{
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(401) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	defer rows.Close()
	defer conn.Close();

}