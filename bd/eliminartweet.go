package bd

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func EliminarTweet(id string) error {

	cxt, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	bd := MongoCN.Database("go")
	col := bd.Collection("tweet")

	Object, _ := primitive.ObjectIDFromHex(id)

	condicion := bson.M{"_id": bson.M{"$eq": Object}}

	_, err := col.DeleteOne(cxt, condicion)

	return err

}
