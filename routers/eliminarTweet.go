package routers

import (
	"github.com/IanDex/twitter/db"
	"net/http"
)

func EliminarTweet(w http.ResponseWriter, r *http.Request)  {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "id required", 400)
		return
	}

	err := db.BorroTweet(ID, IDUsuario)
	if err != nil {
		http.Error(w, "Error al eliminar", 400)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
