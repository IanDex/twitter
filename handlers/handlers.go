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
	router.HandleFunc("/verperfil", middleware.CheckDB(middleware.ValidoJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarPerfil", middleware.CheckDB(middleware.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/tweet", middleware.CheckDB(middleware.ValidoJWT(routers.GraboTweet))).Methods("POST")
	router.HandleFunc("/leoTweets", middleware.CheckDB(middleware.ValidoJWT(routers.LeoTweets))).Methods("GET")
	router.HandleFunc("/eliminarTweet", middleware.CheckDB(middleware.ValidoJWT(routers.EliminarTweet))).Methods(http.MethodDelete)

	router.HandleFunc("/uploadImage", middleware.CheckDB(middleware.ValidoJWT(routers.SubirImagen))).Methods(http.MethodPost)
	router.HandleFunc("/getImage", middleware.CheckDB(middleware.ValidoJWT(routers.GetImage))).Methods(http.MethodGet)

	router.HandleFunc("/altaRelacion", middleware.CheckDB(middleware.ValidoJWT(routers.AltaRelacion))).Methods(http.MethodPost)
	router.HandleFunc("/eliminarRelacion", middleware.CheckDB(middleware.ValidoJWT(routers.EliminarRelacion))).Methods(http.MethodDelete)
	router.HandleFunc("/consultarRelacion", middleware.CheckDB(middleware.ValidoJWT(routers.ConsultarRelacion))).Methods(http.MethodGet)

	router.HandleFunc("/listaUsuarios", middleware.CheckDB(middleware.ValidoJWT(routers.VistaUsuario))).Methods(http.MethodGet)
	router.HandleFunc("/leoTweetsSeguidores", middleware.CheckDB(middleware.ValidoJWT(routers.LeoTweetsSeguidores))).Methods(http.MethodGet)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
