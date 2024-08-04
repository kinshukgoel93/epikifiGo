package connection

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func ConnectToMongo() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb+srv://epikfigo:%3Cpassword%3E@epikfigo.kyfktud.mongodb.net/")

	userName := os.Getenv("MONGO_DB_USERNAME")
	password := os.Getenv("MONGO_DB_PASSWORD")

	clientOptions.SetAuth(options.Credential{
		Username: userName,
		Password: password,
	})

	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	log.Println("Connected to mongoDB")

	return client, nil

}

func GetCollectionPointer() *mongo.Collection {
	return collection
}
