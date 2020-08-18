package routers

import (
	"encoding/json"
	"github.com/IanDex/twitter/db"
	"net/http"
)

func VerPerfil(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el ID ", http.StatusBadRequest)
		return
	}

	perfil, err := db.BuscoPerfil(ID)

	if err != nil {
		http.Error(w, "Error al buscar " + err.Error(), 400)
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(perfil)
}
