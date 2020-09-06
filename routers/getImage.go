package routers

import (
	"github.com/IanDex/twitter/db"
	"io"
	"net/http"
	"os"
)

func GetImage(w http.ResponseWriter, r *http.Request)  {
	ID := r.URL.Query().Get("id")
	option := r.URL.Query().Get("opc")
	if len(ID) < 1 {
		http.Error(w, "ID required", 400)
		return
	}

	perfil, err := db.BuscoPerfil(ID)

	if err != nil {
		http.Error(w, "User not found", 400)
		return
	}

	var ruta string = "uploads/avatars/"
	
	if option == "banner" {
		ruta = "uploads/banner/"
	}
	OpenFile, err := os.Open(ruta + perfil.Avatar)

	if err != nil {
		http.Error(w, "Image not found", 400)
		return
	}

	_, err = io.Copy(w, OpenFile)
	if err != nil {
		http.Error(w, "Image not copy", 400)
	}
}
