package jwt

import (
	"time"

	"github.com/Anitori/Project-Go/models"
	jwt "github.com/dgrijalva/jwt-go"
)

func GeneroJWT(t models.Usuario) (string, error){

	//Clave Privada
	miClave := []byte("abcdefg")

	//Claims
	payload := jwt.MapClaims{
		"email": t.Email,
		"nombre": t.Nombre,
		"apellidos": t.Apellidos,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografia": t.Biografia,
		"ubicacion": t.SitioWeb,
		"sitioweb": t.SitioWeb,
		"_id": t.ID.Hex(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}


	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload) //Objeto token
	tokenStr, err := token.SignedString(miClave)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}