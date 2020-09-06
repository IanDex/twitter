package routers

import (
	"github.com/IanDex/twitter/db"
	"github.com/IanDex/twitter/models"
	"net/http"
)

func EliminarRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID
	if len(ID) < 1 {
		http.Error(w, "id required", 400)
		return
	}

	status, err := db.BorroRelacion(t)
	if err != nil && !status {
		http.Error(w, "Error al eliminar", 400)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
