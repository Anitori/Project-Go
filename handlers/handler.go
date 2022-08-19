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


	 PORT := os.Getenv("PORT")
	 if PORT == "" {
		PORT = "8080"
	 }

	 handler := cors.AllowAll().Handler(router)
	 log.Fatal(http.ListenAndServe(":"+PORT,handler))


}