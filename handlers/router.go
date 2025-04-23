package handlers

import (
	"epifigo/services"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Application struct {
	Models services.Models
}

func CreateRouter(app Application) http.Handler {
	log.Println("Hello Router")
	r := mux.NewRouter()
	r.HandleFunc("/healthcheck", app.healthCheck).Methods("GET")
	r.HandleFunc("/createUser", app.CreateUser).Methods("POST")
	r.HandleFunc("/findUser", app.FindUser).Methods("GET")

	return r
}
