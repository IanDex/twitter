package routers

import (
	"encoding/json"
	"net/http"

	"github.com/IanDex/twitter/db"
	"github.com/IanDex/twitter/models"
)

/*Registro s*/

func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Users
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	if len(t.Email) == 0 || len(t.Password) < 6 {
		http.Error(w, err.Error(), 401)
		return
	}

	_, encontrado, _ := db.CheckDuplicateUser(t.Email)
	if encontrado == true {
		http.Error(w, "Duplicated", 400)
		return
	}

	_, status, err := db.InsertUser(t)
	if err != nil {
		http.Error(w, "Error perras "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "Muchos if's", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
