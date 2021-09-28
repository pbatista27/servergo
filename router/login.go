package router

import (
	"encoding/json"
	"net/http"

	"github.com/pbatista27/servergo/bd"
	"github.com/pbatista27/servergo/jwt"
	"github.com/pbatista27/servergo/model"
)

func Login(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("content-type", "application/json")

	var t model.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "usuario y/o  contraseña invalido "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "correo electronico es obligatorio", 400)
		return
	}

	if len(t.Password) == 0 {
		http.Error(w, "el password es obligatorio", 400)
		return
	}

	documento, existe := bd.ChequeoLogin(t.Email, t.Password)

	if existe == false {
		http.Error(w, "usuario o clave incorrecto ", 400)
		return
	}

	jwtkey, err := jwt.GeneroJWT(documento)

	if err != nil {
		http.Error(w, "ocurrio un error al tratar de gernerar el token", 400)
		return
	}

	resp := model.RepuestaLogin{
		Token: jwtkey,
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

}
