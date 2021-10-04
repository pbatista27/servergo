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
	route.HandleFunc("/editar-perfil", middlew.ChequeoBD(middlew.ValidarJWT(router.EditarPerfil))).Methods("PUT")
	route.HandleFunc("/tweet", middlew.ChequeoBD(middlew.ValidarJWT(router.RegistrarTweet))).Methods("POST")
	route.HandleFunc("/tweet", middlew.ChequeoBD(middlew.ValidarJWT(router.LeerTweet))).Methods("GET")
	route.HandleFunc("/tweet", middlew.ChequeoBD(middlew.ValidarJWT(router.EliminarTweet))).Methods("DELETE")

	route.HandleFunc("/subir-avatar", middlew.ChequeoBD(middlew.ValidarJWT(router.SubirAvatar))).Methods("POST")
	route.HandleFunc("/obtener-avatar", middlew.ChequeoBD(middlew.ValidarJWT(router.ObtenerAvatar))).Methods("GET")
	route.HandleFunc("/subir-banner", middlew.ChequeoBD(middlew.ValidarJWT(router.SubirBanner))).Methods("POST")
	route.HandleFunc("/obtener-banner", middlew.ChequeoBD(middlew.ValidarJWT(router.ObtenerBanner))).Methods("GET")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(route)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
