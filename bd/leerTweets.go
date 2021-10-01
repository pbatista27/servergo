package bd

import (
	"context"
	"log"
	"time"

	"github.com/pbatista27/servergo/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func LeerTweets(id string, pagina int64) ([]*model.DevuelvoTweets, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	bd := MongoCN.Database("go")
	col := bd.Collection("tweet")

	var tweets []*model.DevuelvoTweets

	condicion := bson.M{"userid": id}

	opciones := options.Find()
	opciones.SetLimit(20)
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}})
	opciones.SetSkip((pagina - 1) * 20)

	cursor, err := col.Find(ctx, condicion, opciones)

	if err != nil {
		log.Fatal(err.Error())
		return tweets, false
	}

	for cursor.Next(context.TODO()) {
		var tweet model.DevuelvoTweets

		err = cursor.Decode(&tweet)

		if err != nil {
			return tweets, false
		}

		tweets = append(tweets, &tweet)
	}

	return tweets, true

}
