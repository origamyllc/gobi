package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	login_dao "gobi/src/dao/Login"
    JwtToken "gobi/src/dao/Login"
	"net/http"
)

func AuthenticateUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var data login_dao.LoginParams
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}

	fmt.Print(data)
	// 1.1  check if username existts in db
	//      if the username exists get the corresponding password
	//      otherwise throw an error

	// 1.2  if the given password and the expected password are
	//      equal sign the payload with the secret to generate the
	//      token


	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": data.Username,
		"password": data.Password,
	});

	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		fmt.Println(error)
	}

	json.NewEncoder(w).Encode(JwtToken.JwtToken{Token: tokenString})


}
