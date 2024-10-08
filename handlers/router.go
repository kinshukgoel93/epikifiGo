package handlers

import (
	"fmt"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func CreateRouter() *chi.Mux {

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTION"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CRSF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	router.Route("/api", func(router chi.Router) {

		// version 1
		router.Route("/v1", func(router chi.Router) {
			fmt.Println("came to router")
			router.Get("/healthcheck", healthCheck)

			//USERS
			router.Post("/createUser", CreateUser)
			router.Get("/findUser", findUser)

		})
	})

	return router

}
