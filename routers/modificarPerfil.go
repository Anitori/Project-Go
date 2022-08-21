package routers

import (
	"encoding/json"
	"net/http"
	"github.com/Anitori/Project-Go/bd"
	"github.com/Anitori/Project-Go/models"
)

//Endpoint
func ModificarPerfil(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario //Creo un modelo Usuario

	err := json.NewDecoder(r.Body).Decode(&t) //Grabo el body, lo decodifico en el modelo usuario 

	//Si hubo algun error muestro un mensaje, puede que el body vino vacío o que vino con errores de formato, un JSON mal construido
	if err != nil {
		http.Error(w, "Datos Incorrectos "+err.Error(), 400)
		return  //Pongo el return para que no continue y me afecte otra rutina
	}


	//Modifico registro devuelve un booleano y un error
	var status bool

	status, err = bd.ModificoRegistro(t, IDUsuario) //Va mi modelo de usuario que acabo de decodificar del body, y ademas IDUsuario lo vengo capturando con el middleware de validoJWT 

	if err != nil {
		http.Error(w, "Ocurrió un error al intentar modificar el registro. Reintente nuevamente "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado modificar el registro del usuario ", 400)
		return
	}

		w.WriteHeader(http.StatusCreated)

}