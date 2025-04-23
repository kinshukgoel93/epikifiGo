package handlers

import (
	"encoding/json"
	"epifigo/models"
	"log"
	"net/http"
)

type UserResponse struct {
	Name        string `json:"name,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	Email       string `json:"email,omitempty"`
	Msg         string `json:"msg"`
	LoginCode   string `json:"loginCode,omitempty"`
	Code        int    `json:"code"`
}

// CreateUser handles creating a new user
func (app *Application) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	// Decode JSON into user struct
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println("Failed to parse request body:", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Insert user using service layer
	if err := app.Models.UserService.InsertUser(&user); err != nil {
		log.Println("User insert failed:", err)
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"msg":  "User created successfully!",
		"code": 200,
	})
}

// FindUser handles finding a user by phone number and validating login code
func (app *Application) FindUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello User")
	phoneNumber := r.URL.Query().Get("phoneNumber")
	loginCode := r.URL.Query().Get("loginCode")

	user, err := app.Models.UserService.FindUserByPhoneNumber(phoneNumber)
	if err != nil {
		log.Println("User not found:", err)
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if user.LoginCode == loginCode {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"name":  user.Name,
			"email": user.Email,
			"msg":   "Login successful!",
			"code":  200,
		})
	} else {
		http.Error(w, "Invalid login code", http.StatusUnauthorized)
	}
}
