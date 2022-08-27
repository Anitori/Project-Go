package models 

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type DevuelvoTweetsSeguidores struct {
	ID 						primitive.ObjectID `bson:"_id" json:"_id,omitempty"` //Todo lo que tenemos en formato JSON es lo que vamos a enviar al http
	UsuarioID 				string `bson:"usuarioid" json:"userId,omitempty"`
	UsuarioRelacionID 		string `bson:"usuariorelacionid" json:"userRelationId,omitempty"`

	Tweet struct {

			Mensaje string	`bson:"mensaje" json:"mensaje,omitempty"`
			Fecha time.Time `bson:"fecha" json:"fecha,omitempty"`
			ID string	`bson:"_id" json:"_id,omitempty"`
	} 
}