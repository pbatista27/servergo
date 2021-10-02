package router

import (
	"encoding/json"
	"net/http"

	"github.com/pbatista27/servergo/bd"
)

func EliminarTweet(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	if len(id) < 1 {
		http.Error(w, "error el identificador de tweet es obligatorio", http.StatusBadRequest)
		return

	}

	err := bd.EliminarTweet(id)

	if err != nil {
		http.Error(w, "ocurrio un error "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(true)

}
