package services

import (
	"context"
	"time"

	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

type User struct {
	Name        string    `json:"name" bson:"user_name"`
	PhoneNumber int       `json:"phoneNumber" bson:"user_phoneNumber"`
	Email       string    `json:"email" bson:"user_email"`
	CreatedAt   time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

func New(mongo *mongo.Client) User {
	client = mongo

	return User{}
}

func returnCollectionPointer(collection string) *mongo.Collection {
	return client.Database("epikifigo").Collection(collection)
}

func (t *User) InsertUser(user User) error {
	collection := returnCollectionPointer("users")
	_, err := collection.InsertOne(context.TODO(), User{
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})
	if err != nil {
		log.Println("Inserting Error:", err)
		return err
	}

	return nil
}
