package router

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/pbatista27/servergo/bd"
)

func LeerTweet(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	pagina := r.URL.Query().Get("pagina")

	if len(ID) < 1 {
		http.Error(w, "error debes de enviarnos un id ", http.StatusBadRequest)
		return
	}

	if len(pagina) < 1 {
		http.Error(w, "error debes de enviarnos el parametro pagina ", http.StatusBadRequest)
		return
	}
	pagina2, err := strconv.ParseInt(pagina, 10, 64)

	if err != nil {
		http.Error(w, "error el paramentro pagina debe ser un numero entero mayor a cero", http.StatusBadRequest)
		return
	}

	tweets, verda := bd.LeerTweets(ID, pagina2)

	if verda == false {
		http.Error(w, "error al trata de leer los tweets "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tweets)

}
