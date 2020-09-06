package routers

import (
	"github.com/IanDex/twitter/db"
	"github.com/IanDex/twitter/models"
	"net/http"
)

func AltaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "ID require", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := db.InsertoRelacion(t)
	if err != nil {
		http.Error(w, "error insert", http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "Error grave "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
