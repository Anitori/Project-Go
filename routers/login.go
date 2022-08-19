package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Anitori/Project-Go/bd"
	"github.com/Anitori/Project-Go/jwt"
	"github.com/Anitori/Project-Go/models"
)

func Login(w http.ResponseWriter, r *http.Request){
	w.Header().Add("content-type", "application/json")

	var t models.Usuario

	err:= json.NewDecoder(r.Body).Decode(&t)

	if err !=nil{
		http.Error(w, "Usuario y/o contraseña invalidas "+err.Error(), 400)
		return
	}

	if len(t.Email)==0 {
		http.Error(w, "El email del usuario es requerido ", 400)
		return
	}

	documento, existe :=bd.IntentoLogin(t.Email, t.Password)

	if existe == false {
		http.Error(w, "Usuario y/o contraseña invalidas ", 400)
		return
	}

	jwtKey, err := jwt.GeneroJWT(documento)  //Recibira el documento y me va a devolver el token en string
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar generar el Token Correspondiente"+err.Error(), 400)
		return
	}

	resp := models.RespuestaLogin {
		Token : jwtKey,        //Un token es un acceso que permite la propagacíon de identidad y privilegios 
	}


	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)


	//Como grabar una cookie desde el backend
	//Hacer que el jwt se grabe en la cookie del usuario

	//expirationTime := time.Now().add(24 * time.Hour)           
	//http.SetCookie (w, &http.Cookie{
		//Name: "token",
		//Value: jwtKey,
		//Expires: expirationTime,
	//})
}