package bd

import (
	"github.com/gabriel70g/twittor/models"
	"golang.org/x/crypto/bcrypt"
)

/*IntentoLogin Intento login*/
func IntentoLogin(email string, password string) (models.Usuario, bool) {
	usu, encontrado, _ := ChequeoYaExisteUsuario(email)
	if !encontrado {
		return usu, false
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(usu.Password)

	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return usu, false
	}
	return usu, true
}
