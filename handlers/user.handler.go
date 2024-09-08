package handlers

import (
	"encoding/json"
	"epifigo/services"
	"fmt"
	"log"
	"net/http"
)

var user services.User

type UserResponse struct {
	Name        string
	PhoneNumber string
	Email       string
	Msg         string
	LoginCode   string
	Code        int
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal("Body error", err)
	}

	err = user.InsertUser(user)
	if err != nil {
		errorRes := UserResponse{
			Msg:  "Error",
			Code: 304,
		}
		fmt.Println("Kinshuk err", err)
		json.NewEncoder(w).Encode(errorRes)
		return
	}

	res := UserResponse{
		Msg:  "We got you covered ! Check for unique code over mail.",
		Code: 200,
	}

	jsonStr, err := json.Marshal(res)
	if err != nil {
		log.Fatal("Response", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Code)
	w.Write(jsonStr)
}

func findUser(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Finding User Please wait....")
	phoneNumber := r.URL.Query().Get("phoneNumber")
	loginCode := r.URL.Query().Get("loginCode")
	fmt.Println("phone number", phoneNumber)

	userResponse, err := user.FindUserByPhoneNumber(phoneNumber)

	if err != nil {
		errorRes := UserResponse{
			Msg:  "Error",
			Code: 304,
		}
		json.NewEncoder(w).Encode(errorRes)
		return
	}
	var res UserResponse
	if userResponse.LoginCode == loginCode {
		res = UserResponse{
			Name:        userResponse.Name,
			PhoneNumber: userResponse.PhoneNumber,
			Email:       userResponse.Email,
			Code:        200,

			Msg: "Welcome " + userResponse.Name,
		}
	} else {
		res = UserResponse{
			Msg:  "Please find the correct login code",
			Code: 404,
		}
	}

	jsonStr, err := json.Marshal(res)
	if err != nil {
		log.Fatal("Response", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(res.Code)
	w.Write(jsonStr)
}
