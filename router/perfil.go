package router

import (
	"encoding/json"
	"net/http"

	"github.com/pbatista27/servergo/bd"
)

func Perfil(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "debe enviar el id", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscarPerfil(ID)
	if err != nil {
		http.Error(w, "ocurrio un error en el registro "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)
}
