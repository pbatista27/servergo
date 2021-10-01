package router

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/pbatista27/servergo/bd"
	"github.com/pbatista27/servergo/model"
)

func RegistrarTweet(w http.ResponseWriter, r *http.Request) {

	var t model.Twett

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "ocurrio un error al recibir los datos "+err.Error(), http.StatusBadRequest)
		return
	}

	registro := model.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: t.Mensaje,
		Fecha:   time.Now(),
	}

	var statu bool

	_, statu, err = bd.InsertarTweet(registro)

	if err != nil {
		http.Error(w, "ocurrio un error al registrar tweet "+err.Error(), http.StatusBadRequest)
		return
	}

	if statu == false {
		http.Error(w, "ocurrio un error al registrar tweet", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
