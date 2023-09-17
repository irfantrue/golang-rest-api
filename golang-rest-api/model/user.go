package model

import (
	"golang-rest-api/config/mongodb"

	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type PayloadLogin struct {
	Username string `validate:"required"`
	Email    string `validate:"required,min=5"`
}

func UserCollection() *mongo.Collection {
	mongo, err := mongodb.Connect()
	if err != nil {
		panic(err)
	}
	return mongo.Database("development").Collection("users")
}
