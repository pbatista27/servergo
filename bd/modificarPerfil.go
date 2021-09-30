package bd

import (
	"context"
	"time"

	"github.com/pbatista27/servergo/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ModificarPerfil(u model.Usuario, ID string) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("go")
	col := db.Collection("usuarios")

	registro := make(map[string]interface{})

	registro["nombre"] = u.Nombre
	registro["apellidos"] = u.Apellidos

	updateStr := bson.M{
		"$set": registro,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filtro := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filtro, updateStr)

	if err != nil {
		return false, err
	}

	return true, nil
}
