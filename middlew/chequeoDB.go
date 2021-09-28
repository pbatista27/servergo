package middlew

import (
	"net/http"

	"github.com/pbatista27/servergo/bd"
)

func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "conexion perdida de base datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
