package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gabriel70g/twittor/bd"
	"github.com/gabriel70g/twittor/models"
)

/*Email valor usado en todos los endpoints*/
var Email string

/*IDUsuario ID es el id devuelto del modelo*/
var IDUsuario string

/*ProcesoToken Proceso token para extraer sus valors */
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("La clave para el proyecto de Twittot")
	claims := &models.Claim{}

	splipToken := strings.Split(tk, "Bearer")
	if len(splipToken) != 2 {
		return claims, false, string(""), errors.New("Formato de token invalido")
	}
	tk = strings.TrimSpace((splipToken[1]))

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		_, encontrado, ID := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, ID, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("Formato de token invalido")
	}

	return claims, false, string(""), err
}
