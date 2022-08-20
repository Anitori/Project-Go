package routers

import(
	"encoding/json"
	"net/http"
	"github.com/Anitori/Project-Go/bd"
)

//Permite extraer los valores del perfil
func VerPerfil(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")    //En postman, cuando es un GET y se envían datos, se hace desde parámetros, dichos parámetros se me van a incluir en la URL, de esa URL va a capturar el ID. Se separa la URL de los parámetros con un signo de pregunta.
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar buscar el registro "+err.Error(), 400)
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)

}