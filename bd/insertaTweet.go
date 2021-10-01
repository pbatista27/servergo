package bd

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/pbatista27/servergo/model"
)

func InsertarTweet(t model.GraboTweet) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("go")
	col := db.Collection("tweet")

	registro := bson.M{
		"userid":  t.UserID,
		"mensaje": t.Mensaje,
		"fecha":   t.Fecha,
	}

	result, err := col.InsertOne(ctx, registro)

	if err != nil {
		return "", false, err
	}

	ObjecID, _ := result.InsertedID.(primitive.ObjectID)

	return ObjecID.String(), true, nil
}
