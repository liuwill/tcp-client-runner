package utils

import (
	"gopkg.in/mgo.v2/bson"
)

func GenerateObjectId() string {
	objectId := bson.NewObjectId()
	return objectId.Hex()
}
