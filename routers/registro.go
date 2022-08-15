package routers

import (
	"encoding/json" //Paquete de GO que nos permite trabajar con todo lo que tenga JSON
	"net/http"
	"github.com/Anitori/Project-Go/bd"
	"github.com/Anitori/Project-Go/models"

)

//Registro es la función para crear en la BD el registro de usuario

func Registro(w http.ResponseWriter, r *http.Request){
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t) //El body de un objeto HTTP es un Stream (es un dato que se puede leer una sola vez)
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}


	
//Hay que hacer todas las validaciones correspondientes

if len(t.Email)==0 {  
	http.Error(w, "El email de usuario es requerido", 400)
	return
}

if len(t.Password) < 6 {  
	http.Error(w, "Debe especificar una contraseña de al menos 6 caracteres", 400)
	return
}

// Validación contra los datos

//Cuando en un parametro, no voy a utilizar los otros, pongo un _

_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email) // Si me devuelve un true, quiere decir que se quiere registrar con un correo ya existente
if encontrado == true {
	http.Error(w, "Ya existe un usuario registrado con ese email", 400)
	return
}


_, status, err := bd.InsertoRegistro(t) 

if err != nil {
	http.Error(w, "Ocurrió un error al intentar realizar el registro de usuario"+err.Error(), 400)
	return
}

if status == false {
	http.Error(w, "No se ha logrado insertar el registro del usuario", 400)
	return
}


w.WriteHeader(http.StatusCreated)

}


