package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	statics "gobi/src/constants/app"
	queries "gobi/src/constants/sql"
	login_dao "gobi/src/dao/Login"
	user_dao "gobi/src/dao/User"
    JwtToken "gobi/src/dao/Login"
	db "gobi/src/database/postgres"
	"log"
	"net/http"
	"os"
)

func AuthenticateUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var data login_dao.LoginParams
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}

	// 1.1  check if username existts in db
	//      if the username exists get the corresponding password
	//      otherwise throw an error

	// 1.2  if the given password and the expected password are
	//      equal sign the payload with the secret to generate the
	//      token


	conn, err := db.Connect();
	if err != nil {
		panic(err)
	}
	host, _ := os.Hostname();
	port := statics.PORT;
	sqlStatement := fmt.Sprintf(queries.GET_USER_BY_USER_NAME_AND_PASSWORD_QUERY, data.Username, data.Password )
	var userDao user_dao.User;
	var users []user_dao.User;
	rows,err := db.Get(sqlStatement,conn)
	for rows.Next() {
		error :=  rows.Scan(&userDao.ID, &userDao.Username, &userDao.Password,&userDao.Secret, &userDao.FirstName, &userDao.LastName, &userDao.Email,&userDao.OrganizationName,&userDao.Created,&userDao.Updated)
		if error != nil {
			log.Fatal("Host: "+host +"Port: " +port)
			log.Fatal(error)
			panic(error)
		} else{
			users = append(users,userDao);
		}
	}
	m,err := json.Marshal(users);
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	log.Print("Host: "+host +"Port: " +port+ " result:= "+string(m))
	if string(m) != "null"  {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": data.Username,
			"password": data.Password,
		});

		tokenString, error := token.SignedString([]byte("secret"))
		if error != nil {
			fmt.Println(error)
		}

		json.NewEncoder(w).Encode(JwtToken.JwtToken{Token: tokenString})
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
