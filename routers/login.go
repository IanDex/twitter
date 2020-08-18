package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/IanDex/twitter/jwt"
	"github.com/IanDex/twitter/db"
	"github.com/IanDex/twitter/models"
)

// Login CheckDuplicateUser
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var user models.Users

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Usuario y/k " + err.Error(), 400)
	}

	if len(user.Email) == 0 {
		http.Error(w, "El email del usuario es requerido ", 400)
		return
	}

	documento, existe := db.IntentoLogin(user.Email, user.Password)
	if !existe {
		http.Error(w, "s", 400)
		return
	}

	jwtKey, err := jwt.GeneroJWT(documento)
	if err != nil {
		http.Error(w, "Token error ", 400)
	}

	resp := models.RespuestaLogin{
		Token : jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Value: jwtKey,
		Expires: expirationTime,
	})


}
