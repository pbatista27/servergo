package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN  viriable de conexion */
var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb://localhost:27017/go")

/*ConctarDB funcion de conectar*/
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("conexion exitosa")
	return client
}

/*ChequeoConection chequea el ping de la conexion */
func ChequeoConnection() bool {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {

		return false
	}

	return true
}
