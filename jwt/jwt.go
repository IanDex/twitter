package jwt

import (
	"time"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/IanDex/twitter/models"
)

func GeneroJWT(user models.Users) (string, error)  {
	secretKey := []byte("Cristtian")
	payload := jwt.MapClaims{
		"email": user.Email,
		"nombre": user.Nombre,
		"apellidos" : user.Apellidos,
		"fechaNac": user.FechaNac,
		"biografia": user.Biografia,
		"ubicacion": user.Ubicacion,
		"sitioweb": user.SitioWeb,
		"_id": user.ID.Hex(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(secretKey)
	if err != nil {
		 return "", err
	}
	return tokenStr, nil
}