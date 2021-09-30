package router

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/pbatista27/servergo/bd"
	"github.com/pbatista27/servergo/model"
)

var Email string
var IDUsuario string

func ProcesarJWT(tk string) (*model.Claim, bool, string, error) {

	miClave := []byte("i4ijiijd,.lÑLDjidjij4jdi")
	claims := &model.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err == nil {

		_, encontrado, _ := bd.ChequeYaExisteUsuario(claims.Email)

		if encontrado == true {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}

		return claims, encontrado, IDUsuario, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")

	}
	return claims, false, string(""), err

}
