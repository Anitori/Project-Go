package handlers

import (
	"log"
	"net/http"
	"os"
	"github.com/Anitori/Project-Go/middlew"
	"github.com/Anitori/Project-Go/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//Seteo mi puerto, el handler y pongo a escuchar al servidor

func Manejadores () { // Función que va a ejecutar cuando llame a la API, define cada una de las rutas 
	 router := mux.NewRouter() // Lo que hace el max es capturar el http y llevar el manejo al response writter y al request que tiene en el llamado de la API

	//Rutas
	
	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.VerPerfil))).Methods("GET") // El GET es método de petición, 
	router.HandleFunc("/modificarPerfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.ChequeoBD(middlew.ValidoJWT(routers.GraboTweet))).Methods("POST")
	router.HandleFunc("/leoTweets", middlew.ChequeoBD(middlew.ValidoJWT(routers.LeoTweets))).Methods("GET")
	router.HandleFunc("/eliminarTweet", middlew.ChequeoBD(middlew.ValidoJWT(routers.EliminarTweet))).Methods("DELETE")

	router.HandleFunc("/subirAvatar", middlew.ChequeoBD(middlew.ValidoJWT(routers.SubirAvatar))).Methods("POST") //Cuando Envío es POST
	router.HandleFunc("/obtenerAvatar", middlew.ChequeoBD(middlew.ValidoJWT(routers.ObtenerAvatar))).Methods("GET") //Cuando leo es GET
	router.HandleFunc("/subirBanner", middlew.ChequeoBD(middlew.ValidoJWT(routers.SubirBanner))).Methods("POST") 
	router.HandleFunc("/obtenerBanner", middlew.ChequeoBD(middlew.ValidoJWT(routers.ObtenerBanner))).Methods("GET") 

	router.HandleFunc("/altaRelacion", middlew.ChequeoBD(middlew.ValidoJWT(routers.AltaRelacion))).Methods("POST") 
	router.HandleFunc("/bajaRelacion", middlew.ChequeoBD(middlew.ValidoJWT(routers.BajaRelacion))).Methods("DELETE") 



	 PORT := os.Getenv("PORT")
	 if PORT == "" {
		PORT = "8080"
	 }

	 handler := cors.AllowAll().Handler(router)
	 log.Fatal(http.ListenAndServe(":"+PORT,handler))


}