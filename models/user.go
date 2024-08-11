package models

import "time"

type User struct {
	Name              string    `json:"name" bson:"user_name"`
	PhoneNumber       int       `json:"phoneNumber" bson:"user_phoneNumber"`
	YearsOfExperience int       `json:"yearsOfExperience" bson:"user_yearsOfExperience"`
	Email             string    `json:"email" bson:"user_email"`
	CreatedAt         time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt         time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
