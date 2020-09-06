package routers

import (
	"encoding/json"
	"github.com/IanDex/twitter/db"
	"net/http"
	"strconv"
)

func VistaUsuario(w http.ResponseWriter, r *http.Request) {
	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pagTemp, err := strconv.Atoi(page)

	if err != nil {
		http.Error(w, "p", http.StatusBadRequest)
		return
	}

	pag := int64(pagTemp)

	result, status := db.LeoUsuarioTodos(IDUsuario, pag, search, typeUser)

	if !status {
		http.Error(w, "p", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(result)
}
