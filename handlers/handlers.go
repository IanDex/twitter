package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/IanDex/twitter/middleware"
	"github.com/IanDex/twitter/routers"
)

/*Manejadores d*/
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middleware.CheckDB(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middleware.CheckDB(routers.Login)).Methods("POST")
	//router.HandleFunc("/verperfil", middleware.ValidoJWT(routers.VerPerfil)).Methods("POST")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
