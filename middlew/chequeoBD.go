package middlew

import (
	"github.com/Anitori/Project-Go/bd"
	"net/http"
)

//ChequeoBD es el middlew que me permite conocer el estado de la BD

func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConecction() == 0 {
			http.Error(w, "Conexión perdida con la base de datos", 500) 
			return
			//Cuando la API haga el llamado al endpoint del registro y encuentre que hay un problema va a recibir automaticamente un status 500
		}

		next.ServeHTTP(w, r)  //Si da error, me manda un return, y si no, esta devolviendo una función donde estoy pasando al próximo eslabon de la cadena todos los objetos de ResponseWritter y request
	}
}
