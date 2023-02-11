package util

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetRandomID() string {
	return primitive.NewObjectID().Hex()
}
