package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	queries "gobi/src/constants/sql"
	statics "gobi/src/constants/app"
	db "gobi/src/database/postgres"
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
	 	fmt.Print(rows);
	 }
	defer conn.Close();
}

