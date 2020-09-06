package routers

import (
	"github.com/IanDex/twitter/db"
	"github.com/IanDex/twitter/models"
	"io"
	"net/http"
	"os"
	"strings"
)

func SubirImagen(w http.ResponseWriter, r *http.Request)  {
	file, _handler, err := r.FormFile("avatar")
	option := r.Form.Get("opc")
	var usuario models.Users
	var extension = strings.Split(_handler.Filename, ".")[1]
	var nameFile string = IDUsuario + "." + extension
	var archivo string = "uploads/avatars/" + nameFile

	if option == "banner" {
		usuario.Banner = nameFile
		archivo = "uploads/banner/" + nameFile
	}else {
		usuario.Avatar = nameFile
	}

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error upload", 400)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error copy", 400)
		return
	}

	var status bool

	status, err = db.ModificoRegistro(usuario, IDUsuario)
	if err != nil || !status {
		http.Error(w, "Error al grabar avatar", 400)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)


}