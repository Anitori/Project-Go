package routers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/Anitori/Project-Go/bd"
)


func LeoTweets(w http.ResponseWriter, r *http.Request){
	ID := r.URL.Query().Get("id")
	if len(ID)< 1{
		http.Error(w, "Debe enviar el parámetro id", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("pagina"))<1 {
		http.Error(w, "Debe enviar el parámetro página", http.StatusBadRequest)
		return
	}

	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))  //Conversion de un string a entero

	if err != nil {
		http.Error(w, "Debe enviar el parametro página con un valor mayor a 0", http.StatusBadRequest)
		return
	}

	pag := int64(pagina) //La rutina que usamos para paginar en bson requiere que el dato sea int64
	respuesta, correcto := bd.LeoTweets(ID, pag)
	if correcto == false {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)

}