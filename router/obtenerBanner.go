package router

import (
	"io"
	"net/http"
	"os"

	"github.com/pbatista27/servergo/bd"
)

func ObtenerBanner(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debes enviar el id", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscarPerfil(ID)

	if err != nil {
		http.Error(w, "usuario no encontrado "+err.Error(), http.StatusBadRequest)
		return
	}

	OpenFile, err := os.Open("uploads/banners/" + perfil.Avatar)

	if err != nil {
		http.Error(w, "Banner no encotrado", http.StatusBadRequest)
	}

	_, err = io.Copy(w, OpenFile)

	if err != nil {
		http.Error(w, "Error al copia img", http.StatusBadRequest)
		return
	}

}
