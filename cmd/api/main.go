package main

import (
	"context"
	"epifigo/connection"
	"epifigo/handlers"
	"epifigo/services"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Application struct {
	Models services.Models
}

func main() {
	mongoClient, err := connection.ConnectToMongo()
	if err != nil {
		fmt.Println("error while connecting to mongo", err)
		log.Panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	defer func() {
		if err = mongoClient.Disconnect(ctx); err != nil {
			fmt.Println("error while disconnecting to mongo", err)
			panic(err)
		}
	}()
	services.New(mongoClient)
	log.Println("Server running in port", 3030)
	connectionError := http.ListenAndServe(":3030", handlers.CreateRouter())

	if errors.Is(connectionError, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
