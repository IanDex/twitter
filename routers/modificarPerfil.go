package routers

import (
	"encoding/json"
	"github.com/IanDex/twitter/db"
	"github.com/IanDex/twitter/models"
	"net/http"
)

func ModificarPerfil(w http.ResponseWriter, r *http.Request){
	var t models.Users

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Datos Incorrectos " + err.Error(), 400)
		return
	}
	var status bool
	status, err = db.ModificoRegistro(t, IDUsuario)
	if err != nil {
		http.Error(w, "Error Modificar Registro " + err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "Registro no modificado", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}