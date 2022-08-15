package bd

import (

	"golang.org/x/crypto/bcrypt"
)


//Esto me permite encriptar la password recibida
func EncriptarPassword(pass string) (string, error){
	costo := 8 

	//Es el número (Es un algoritmo basado en 2 elevado al costo),  El algoritmo GenerateFromPassword va a tomar el número de costo y lo va a llevar al cuadrado, y esa es la cantidad de pasadas que va hacer para encriptarlo. (256 pasadas en este caso) 256 veces encriptado
	
	// Se recomienda que si lo que encriptamos es la contraseña de un usuario común, se coloque un costo de 6, pero si lo que estamos encriptando es el usuario de un superusuario o admin, usemos uno de 8. Si colocamos un valor de costo mayor, puede tomar una gran cantidad de segundos en encriptar una password. 


	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)  //Esta función solo acepta un slice de bytes (Slice: Vector sin cantidad de elementos)
	return string(bytes), err

	
}