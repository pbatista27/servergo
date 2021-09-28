package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/pbatista27/servergo/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BuscarPerfil(ID string) (model.Usuario, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("go")
	col := db.Collection("usuarios")

	var perfil model.Usuario

	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{"_id": objID}

	err := col.FindOne(ctx, condicion).Decode(&perfil)

	perfil.Password = ""

	if err != nil {
		fmt.Println("registro no encontrado" + err.Error())
		return perfil, err
	}

	return perfil, nil

}
