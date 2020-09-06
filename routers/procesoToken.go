package routers

import (
	"errors"
	"github.com/IanDex/twitter/db"
	"github.com/IanDex/twitter/models"
	jwt "github.com/dgrijalva/jwt-go"
	"strings"
)

var Email     string
var	IDUsuario string


func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	secretKey := []byte("Cristtian")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, " ")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("Puto tutorial de mierda")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error){
		return secretKey, nil
	})

	if err == nil {
		_, encontrado, _ := db.CheckDuplicateUser(claims.Email)
		if encontrado {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}

		return claims, encontrado, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, string("err"), errors.New("token invalido")
	}

	return claims, false, string(""), err
}
