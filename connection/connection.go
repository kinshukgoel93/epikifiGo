package connection

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToMongo() (*mongo.Client, error) {
	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb+srv://epikfigo:epikfigo@epikfigo.kyfktud.mongodb.net/?retryWrites=true&w=majority&appName=epikfiGo").SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.Background(), opts)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	dbNames, error := client.ListDatabaseNames(context.Background(), bson.M{})
	if error != nil {
		log.Fatal("error", error)
		return nil, error
	} else {
		fmt.Println("dbNames", dbNames)
	}
	return client, nil
}
