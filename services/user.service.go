package services

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"epifigo/models"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	client     *mongo.Client
	dbName     string
	collection string
}

func NewUserService(client *mongo.Client) *UserService {
	return &UserService{
		client:     client,
		dbName:     "epikifigo",
		collection: "users",
	}
}

func (s *UserService) getCollection() *mongo.Collection {
	return s.client.Database(s.dbName).Collection(s.collection)
}

func generateLoginCode() string {
	token := make([]byte, 4)
	_, err := rand.Read(token)
	if err != nil {
		log.Println("Failed to generate random token:", err)
		return "0000"
	}
	return hex.EncodeToString(token)
}

func (s *UserService) InsertUser(user *models.User) error {
	user.LoginCode = generateLoginCode()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err := s.getCollection().InsertOne(context.TODO(), user)
	if err != nil {
		log.Println("Error inserting user:", err)
		return err
	}

	fmt.Println("Inserted user with login code:", user.LoginCode)
	return nil
}

func (s *UserService) FindUserByPhoneNumber(phone string) (*models.User, error) {
	log.Println("Finding user with phone:", phone)

	var user models.User
	filter := bson.M{"user_phoneNumber": phone}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := s.getCollection().FindOne(ctx, filter).Decode(&user)
	if err != nil {
		log.Println("Error finding user:", err)
		return nil, err
	}
	log.Println("Finding user with phone:", &user)

	return &user, nil
}
