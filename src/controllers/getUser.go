package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	statics "gobi/src/constants/app"
	queries "gobi/src/constants/sql"
	user_dao "gobi/src/dao/User"
	db "gobi/src/database/postgres"
	"log"
	"net/http"
	"os"
)
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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
	log.Print("Host: "+host +"Port: " +port+ " result:= "+string(m))
	json.NewEncoder(w).Encode(	map[string]string{"code": "200", "result": string(m)})
	defer rows.Close()
	defer conn.Close();

}