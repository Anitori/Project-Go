package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Creando el modelo de USUARIOS
//Usuario es el modelo de usuario de la base de MongoDB

type Usuario struct  {
	ID  			primitive.ObjectID 		`bson:"_id,omitempty" json:"id"`  //_id es una convención de mongo
	Nombre 			string 					`bson:"nombre" json:"nombre,omitempty"`
	Apellidos		string 					`bson:"apellidos" json:"apellidos,omitempty"`
	FechaNacimiento	time.Time 				`bson:"fechaNacimiento" json:"fechaNacimiento,omitempty"`
	Email			string 					`bson:"email" json:"email"` //Sacando el omitempty, siempre me va a devolver el Email
	Password		string 					`bson:"password" json:"password,omitempty"` // Siempre va con omitempty, nunca se devuelve una pass por el navegador. Es mala práctica tener un endpoint que me devuelva una password por el navegador. Para todo lo que tiene que ver con credenciales y sobre todo con el login y manejo de sesión vamos a manejar el token. 
	Avatar			string 					`bson:"avatar" json:"avatar,omitempty"`
	Banner			string 					`bson:"banner" json:"banner,omitempty"`
	Biografia		string 					`bson:"biografia" json:"biografia,omitempty"`
	Ubicacion		string 					`bson:"ubicacion" json:"ubicacion,omitempty"`
	SitioWeb		string 					`bson:"sitioweb" json:"sitioWeb8,omitempty"`

}