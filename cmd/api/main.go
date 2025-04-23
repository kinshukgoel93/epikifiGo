package main

import (
	"context"
	"epifigo/connection"
	"epifigo/handlers"
	"epifigo/services"
	"errors"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	mongoClient, err := connection.ConnectToMongo()
	if err != nil {
		log.Panicf("MongoDB connection error: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	defer func() {
		if err := mongoClient.Disconnect(ctx); err != nil {
			log.Panicf("MongoDB disconnect error: %v", err)
		}
	}()

	// Initialize service layer
	models := services.NewModels(mongoClient)
	app := handlers.Application{Models: models}

	log.Println("✅ Server running on port 3030")
	if err := http.ListenAndServe(":3030", handlers.CreateRouter(app)); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			log.Println("Server closed")
		} else {
			log.Fatalf("Server error: %v", err)
			os.Exit(1)
		}
	}
}
