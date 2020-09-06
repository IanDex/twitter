package routers

import (
	"encoding/json"
	"github.com/IanDex/twitter/db"
	"github.com/IanDex/twitter/models"
	"net/http"
)

func ConsultarRelacion(w http.ResponseWriter, r *http.Request)  {
	ID := r.URL.Query().Get("id")
	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	var resp models.RespuestaC_Relacion

	status, err := db.ConsultaRelacion(t)
	resp.Status = true
	if err != nil || !status {
		resp.Status = false
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(resp)
}