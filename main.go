package main

import (
	"log"

	"github.com/pbatista27/servergo/bd"
	"github.com/pbatista27/servergo/handler"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("sin conexion")
		return
	}
	handler.Manejadores()
}
