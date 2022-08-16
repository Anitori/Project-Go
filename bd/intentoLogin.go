package bd 

import (
	"github.com/Anitori/Project-Go/models"
	"golang.org/x/crypto/bcrypt"
)


//Realiza el chequeo de login a la BD
func IntentoLogin(email string, password string) (models.Usuario, bool) {

	usu, encontrado, _ := ChequeoYaExisteUsuario(email)
	if encontrado == false {
		return usu, false
	}

	//Comparar la password 

	//Grabar en 2 variables la password que me vino de parámetro en la función, y la que esta grabada en la base de datos 
	passwordBytes := []byte(password)
	passwordBD := []byte(usu.Password)

	//Hash es la encriptada, y la password se refiere a la original
	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)   
	if err != nil {
		return usu, false
	}

	return usu, true

}