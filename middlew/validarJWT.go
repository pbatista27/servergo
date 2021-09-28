package middlew

import (
	"net/http"

	"github.com/pbatista27/servergo/router"
)

func ValidarJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := router.ProcesarJWT(r.Header.Get("authorization"))

		if err != nil {
			http.Error(w, "error en el token!"+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}
