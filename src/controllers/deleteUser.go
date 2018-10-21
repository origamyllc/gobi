package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	statics "gobi/src/constants/app"
	queries "gobi/src/constants/sql"
	db "gobi/src/database/postgres"
	"log"
	"net/http"
	"os"
)

func DeleteUser(w http.ResponseWriter, r *http.Request){
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
	 rows,err := db.Get(sqlStatement,conn)
	 if err == nil {
		 log.Print("Host: "+host +"Port: " +port+ " result:= Deleted")
		 w.WriteHeader(http.StatusCreated)
		 json.NewEncoder(w).Encode(map[string]string{"code":"200","message": "record sucessfully deleted"})
	 }else {
		 log.Fatal("Host: "+host +"Port: " +port+ " result:= Not Deleted")
		 log.Fatal(rows)
		 w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		 w.WriteHeader(500) // unprocessable entity
		 if err := json.NewEncoder(w).Encode(err); err != nil {
			 panic(err)
		 }
	 }
	defer conn.Close();
}

