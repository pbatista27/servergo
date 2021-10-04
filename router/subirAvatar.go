package router

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/pbatista27/servergo/bd"
	"github.com/pbatista27/servergo/model"
)

func SubirAvatar(w http.ResponseWriter, r *http.Request) {

	file, handler, err := r.FormFile("avatar")

	var extension = strings.Split(handler.Filename, ".")[1]

	var archivo string = "uploads/avatars/" + IDUsuario + "." + extension

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "error al subir archivo"+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "error al guardar el archivo"+err.Error(), http.StatusBadRequest)
		return
	}

	var usuario model.Usuario
	var status bool

	usuario.Avatar = IDUsuario + "." + extension
	status, err = bd.ModificarPerfil(usuario, IDUsuario)

	if err != nil || status == false {
		http.Error(w, "error al subir archivo del avatar"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
