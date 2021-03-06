package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gabriel70g/twittor/middlew"
	"github.com/gabriel70g/twittor/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Manejadores seteo mi puerto, el Handler y pongo a escuchar al Servidor */
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/", home)
	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlew.ChequeoBD(middlew.ValidJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarperfil", middlew.ChequeoBD(middlew.ValidJWT(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.ChequeoBD(middlew.ValidJWT(routers.GraboTweet))).Methods("POST")
	router.HandleFunc("/leotweets", middlew.ChequeoBD(middlew.ValidJWT(routers.LeoTweets))).Methods("GET")
	router.HandleFunc("/eliminartweet", middlew.ChequeoBD(middlew.ValidJWT(routers.EliminarTweet))).Methods("DELETE")

	router.HandleFunc("/subiravatar", middlew.ChequeoBD(middlew.ValidJWT(routers.SubirAvatar))).Methods("POST")
	router.HandleFunc("/subirbanner", middlew.ChequeoBD(middlew.ValidJWT(routers.SubirBanner))).Methods("POST")
	router.HandleFunc("/obteneravatar", middlew.ChequeoBD(routers.ObtenerAvatar)).Methods("GET")
	router.HandleFunc("/obtenerbanner", middlew.ChequeoBD(routers.ObtenerBanner)).Methods("GET")

	router.HandleFunc("/altarelacion", middlew.ChequeoBD(middlew.ValidJWT(routers.AltaRelacion))).Methods("POST")
	router.HandleFunc("/bajarelacion", middlew.ChequeoBD(middlew.ValidJWT(routers.BajaRalacion))).Methods("DELETE")
	router.HandleFunc("/consultarelacion", middlew.ChequeoBD(middlew.ValidJWT(routers.ConsultaRelacion))).Methods("GET")

	router.HandleFunc("/listarusuarios", middlew.ChequeoBD(middlew.ValidJWT(routers.ListaUsuarios))).Methods("GET")
	router.HandleFunc("/leotweetsseguidores", middlew.ChequeoBD(middlew.ValidJWT(routers.LeoTweetsSeguidores))).Methods("GET")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}
