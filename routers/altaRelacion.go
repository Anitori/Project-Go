package routers

import (
	"net/http"

	"github.com/Anitori/Project-Go/bd"
	"github.com/Anitori/Project-Go/models"
)

//Realiza el registro de la relación entre usuarios
func AltaRelacion(w http.ResponseWriter, r* http.Request){

	ID := r.URL.Query().Get("id")
	if len(ID)<1 {
		http.Error(w, "El parámetro ID es obligatorio", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID


	status, err := bd.InsertoRelacion(t)
	if err != nil {
		http.Error(w, "Ocurrío un error al intentar insertar relación "+err.Error(),http.StatusBadRequest )
		return
	}


	if status == false{
		http.Error(w, "No se ha logrado insertar la relación"+err.Error(),http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}