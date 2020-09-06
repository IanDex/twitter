package routers

import (
	"encoding/json"
	"github.com/IanDex/twitter/db"
	"github.com/IanDex/twitter/models"
	"net/http"
	"time"
)

func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := db.InsertoTweet(registro)

	if err != nil {
		http.Error(w, "Error insertar tweet", 400)
		return
	}

	if !status {
		http.Error(w, "NOpe tweet", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
