package middlew

import (
	"net/http"
	"github.com/Anitori/Project-Go/routers"
)

//Permite validar el JWT que nos viene en la petición

func ValidoJWT(next http.HandlerFunc) http.HandlerFunc {  //Recordar que cuando es un middleware, recibe y devuelve lo mismo
	return func(w http.ResponseWriter, r *http.Request){ //Va a llamar a una rutina creada en routers que va a devolver 4 parámetros (acá solo queremos saber si hay o no error en la ejecución de dicha rutina)

	_, _, _, err := routers.ProcesoToken(r.Header.Get("Authorization"))
	if err != nil {
		http.Error(w, "Error en el Token ! "+err.Error(), http.StatusBadRequest)
	}

	next.ServeHTTP(w, r)
	}
} 