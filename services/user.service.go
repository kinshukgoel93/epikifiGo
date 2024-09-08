package services

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

type User struct {
	Name        string    `json:"name" bson:"user_name"  binding:"required,alpha"`
	PhoneNumber string    `json:"phoneNumber" bson:"user_phoneNumber"`
	Email       string    `json:"email" bson:"user_email" binding:"required,email"`
	CreatedAt   time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	LoginCode   any       `json:"code" bson:"user_code"  binding:"required,alpha"`
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
	token := make([]byte, 4)
	rand.Read(token)
	tokenCode := fmt.Sprintf("%x", token)
	print("tokenCode", tokenCode)
	_, err := collection.InsertOne(context.TODO(), User{
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		LoginCode:   tokenCode,
	})
	if err != nil {
		log.Println("Inserting Error:", err)

		return err
	}

	return nil
}

func (t *User) FindUserByPhoneNumber(phoneNumber string) (User, error) {
	collection := client.Database("epikifigo").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var user User

	// Find the user based on phone number
	filter := bson.M{"user_phoneNumber": phoneNumber}

	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return User{}, err
	}

	return user, nil
}
