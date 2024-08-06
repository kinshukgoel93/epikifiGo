package main

import (
	"context"
	"epifigo/connection"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Welcome to EPIFIGO")
	mongoClient, err := connection.ConnectToMongo()
	if err != nil {
		log.Panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	defer func() {
		if err = mongoClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	log.Println("Server running in port", 8080)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
