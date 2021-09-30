package router

import (
	"encoding/json"
	"net/http"

	"github.com/pbatista27/servergo/bd"
	"github.com/pbatista27/servergo/model"
)

func EditarPerfil(w http.ResponseWriter, r *http.Request) {

	var u model.Usuario

	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		http.Error(w, "ocurrio un error al enviar los datos "+err.Error(), http.StatusBadRequest)
		return
	}

	var existe bool
	existe, err = bd.ModificarPerfil(u, IDUsuario)

	if existe == false {
		http.Error(w, "usuario no encontrado", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, "ocurrio un error de "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
