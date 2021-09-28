package router

import (
	"encoding/json"
	"net/http"

	"github.com/pbatista27/servergo/bd"
	"github.com/pbatista27/servergo/model"
)

func Registro(w http.ResponseWriter, r *http.Request) {

	var t model.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "error en los datos recibidos "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El campo email es requerido", 400)
		return
	}

	if len(t.Password) == 0 {
		http.Error(w, "El campo password es requirido", 400)
		return
	}

	if len(t.Email) < 6 {
		http.Error(w, "La contraseña debe tener almenos 6 caracteres", 400)
		return
	}

	_, encontrado, _ := bd.ChequeYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "Ya existe un usuario registrado con ese email", 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrio un erroe al intenter registrar el usuario "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "no se ha logrado inserta el registro del usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
