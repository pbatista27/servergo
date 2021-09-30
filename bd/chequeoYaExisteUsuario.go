package bd

import (
	"context"
	"time"

	"github.com/pbatista27/servergo/model"
	"go.mongodb.org/mongo-driver/bson"
)

func ChequeYaExisteUsuario(email string) (model.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("go")
	col := db.Collection("usuarios")
	condicion := bson.M{"email": email}

	var resultado model.Usuario

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}

	return resultado, true, ID
}
