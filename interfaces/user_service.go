package interfaces

import "epifigo/models"

type UserService interface {
	InsertUser(user *models.User) error
	FindUserByPhoneNumber(phone string) (*models.User, error)
}
