package handlers

import (
	"encoding/json"
	"epifigo/services"
	"fmt"
	"log"
	"net/http"
)

var user services.User

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("came to create User")
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal("Body error", err)
	}

	err = user.InsertUser(user)
	if err != nil {
		errorRes := Response{
			Msg:  "Error",
			Code: 304,
		}
		fmt.Println("Kinshuk err", err)
		json.NewEncoder(w).Encode(errorRes)
		return
	}

	res := Response{
		Msg:  "Succesfully Created Todo",
		Code: 201,
	}

	jsonStr, err := json.Marshal(res)
	if err != nil {
		log.Fatal("Response", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Code)
	w.Write(jsonStr)
}
