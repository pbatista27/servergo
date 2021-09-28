package handler

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/pbatista27/servergo/middlew"
	"github.com/pbatista27/servergo/router"
	"github.com/rs/cors"
)

func Manejadores() {
	route := mux.NewRouter()

	route.HandleFunc("/registro", middlew.ChequeoBD(router.Registro)).Methods("POST")
	route.HandleFunc("/login", middlew.ChequeoBD(router.Login)).Methods("POST")
	route.HandleFunc("/perfil", middlew.ChequeoBD(middlew.ValidarJWT(router.Perfil))).Methods("GET")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(route)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
