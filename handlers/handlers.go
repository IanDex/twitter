package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/IanDex/twitter/middleware"
	"github.com/IanDex/twitter/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Manejadores d*/
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middleware.CheckDB(routers.Registro)).Methods("POST")

	port := os.Getenv("PORT")
	if port == "" {
		port = "7600"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
