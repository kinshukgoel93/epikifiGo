package services

import (
	"epifigo/interfaces"

	"go.mongodb.org/mongo-driver/mongo"
)

type Models struct {
	UserService interfaces.UserService
}

func NewModels(client *mongo.Client) Models {
	return Models{
		UserService: NewUserService(client),
	}
}
