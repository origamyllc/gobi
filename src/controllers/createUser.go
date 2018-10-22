package controllers

import (
	"encoding/json"
	"fmt"
	 queries "gobi/src/constants/sql"
	user_dao "gobi/src/dao/User"
	JwtToken "gobi/src/dao/Login"
	db "gobi/src/database/postgres"
	statics "gobi/src/constants/app"
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
	xauth := r.Header.Get("x-authorizarion")

	decoder := json.NewDecoder(r.Body)
	var data user_dao.User
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}

	conn, err := db.Connect();
	if err != nil {
		panic(err)
	}
	host, _ := os.Hostname();
	port := statics.PORT;

 query := fmt.Sprintf(queries.ADD_USER_QUERY, data.Username,data.Password,data.Secret, data.FirstName, data.LastName, data.OrganizationName,data.Email )
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
		err = db.Insert(query, conn)
		fmt.Print(err)
		if err == nil {
			log.Print("Host: " + host + "Port: " + port + " result:= Inserted")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(map[string]string{"code": "200", "message": "record sucessfully inserted"})
		} else {
			log.Print("Host: " + host + "Port: " + port + " result:= Not inserted")
			w.WriteHeader(500) // unprocessable entity
			if err := json.NewEncoder(w).Encode(err); err != nil {
				panic(err)
			}
		}
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(401) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	defer rows.Close();
	defer conn.Close();
}

