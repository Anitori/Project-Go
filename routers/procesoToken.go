package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/Anitori/Project-Go/bd"
	"github.com/Anitori/Project-Go/models"
	
)

var Email string //Valor de email usado en todos los EndPoints
var IDUsuario string //Es el ID devuelto del modelo, que se usará en todos los EndPoints

//Proceso Token para extraer sus valores

//Esta rutina, es una de las funciones mas importantes del backend, primero por la cantidad de veces que la vamos a ejecutar, y segundo porque hace la validación mas importante de todas, de procesar el token y ver que la credencial y los privilegios son válidos 
func ProcesoToken(tk string) (*models.Claim, bool, string, error){
	miClave := []byte("abcdefg")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])


	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token)(interface{}, error){
		return miClave, nil
	})

	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado==true {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}

		return claims, encontrado, IDUsuario, nil
	}

	if !tkn.Valid{
		return claims, false, string(""), errors.New("token Inválido")
	}

	return claims, false, string(""), err
}