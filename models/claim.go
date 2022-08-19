package models

import (
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"  //Porque vamos a tener que trabajar con el objet ID
)

//Estructura usada para procesar el JWT

type Claim struct {  //La pongo en mayúsculas porque tengo que exportarla
	Email 	string `json:"email"`
	ID		primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	jwt.StandardClaims 
}

