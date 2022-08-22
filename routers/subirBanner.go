package routers

import (
	"io" //Porque vamos a utilizar operaciones de imput out
	"net/http"  //Porque estamos en un router
	"os" //Porque vamos a ver funciones del sistema operativo
	"strings"
	"github.com/Anitori/Project-Go/bd"
	"github.com/Anitori/Project-Go/models"
)

func SubirBanner(w http.ResponseWriter, r *http.Request){
	file, handler, err := r.FormFile("banner")
	var extension = strings.Split(handler.Filename,".")[1]
	var archivo string = "uploads/banners/" + IDUsuario + "." + extension

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)  // El 0666 son los permisos de lectura, escritura y ejecución

	if err != nil {
		http.Error(w, "Error al subir la imagen! "+err.Error(), http.StatusBadRequest)
		return
	}


	_, err = io.Copy(f, file) 
	if err != nil {
		http.Error(w, "Error al subir el banner !"+err.Error(),http.StatusBadRequest)
		return
	}

	var usuario models.Usuario
	var status bool

	usuario.Banner = IDUsuario + "." + extension
	status, err = bd.ModificoRegistro(usuario, IDUsuario)
	if err != nil || status == false {
		http.Error(w, "Error al grabar el banner en la BD !"+err.Error(),http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}