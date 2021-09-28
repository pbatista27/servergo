package bd

import (
	"github.com/pbatista27/servergo/model"
	"golang.org/x/crypto/bcrypt"
)

func ChequeoLogin(email string, password string) (model.Usuario, bool) {

	usu, encontrado, _ := ChequeYaExisteUsuario(email)

	if encontrado == false {
		return usu, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(usu.Password)

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)

	if err != nil {
		return usu, false
	}
	return usu, true
}
