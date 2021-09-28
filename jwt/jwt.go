package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/pbatista27/servergo/model"
)

func GeneroJWT(t model.Usuario) (string, error) {
	miClave := []byte("i4ijiijd,.lÑLDjidjij4jdi")

	payload := jwt.MapClaims{
		"nombre": t.Nombre,
		"_id":    t.ID.Hex(),
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenStr, err := token.SignedString(miClave)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil

}
