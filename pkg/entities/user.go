package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID           primitive.ObjectID `bson:"_id" json:"id"`
	Login        string             `bson:"login" json:"login"`
	Password     string             `bson:"-" json:"password"`
	PasswordHash string             `bson:"password" json:"-"`
}
