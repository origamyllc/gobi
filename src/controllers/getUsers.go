package controllers

import (
	"encoding/json"
	queries "gobi/src/constants/sql"
	statics "gobi/src/constants/app"
	user_dao "gobi/src/dao/User"
	db "gobi/src/database/postgres"
	"log"
	"net/http"
	"os"
)

// our main function


func GetUsers(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	conn, err := db.Connect();
	if err != nil {
		panic(err)
	}
	host,_ := os.Hostname();
	port := statics.PORT;
	sqlStatement := queries.GET_USERS_QUERY;
	var user user_dao.User;
	var users  [] user_dao.User;
	rows,err := db.Get(sqlStatement,conn)
	for rows.Next() {
		error := rows.Scan(&user.ID, &user.Age, &user.FirstName, &user.LastName, &user.Email)
		if error != nil {
			log.Fatal("Host: "+host +"Port: " +port)
			log.Fatal(error)
			panic(error)
		} else{
			users = append(users,user);
		}
	}
	m,err := json.Marshal(users);
	if string(m) != "null"  {
		log.Print("Host: "+host +"Porsdt: " +port+ " result:= "+string(m))
		json.NewEncoder(w).Encode(map[string]string{"code": "200", "result": string(m)})
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(500) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	defer rows.Close()
	defer conn.Close();

}

