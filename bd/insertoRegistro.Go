package bd 

import (
	"context"
	"time"
	"github.com/Anitori/Project-Go/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertoRegistro es la parada final con la BD para insertar los datos del usuario
func InsertoRegistro(u models.Usuario) (string, bool, error){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel () //Defer es una instrucción que la seteo al comienzo, pero se ejecuta al final de la instrucción
					// el cancel me cancela el timeout


	db := MongoCN.Database("ProjectGo")  //Creo que ese es el nombre de mi base de datos, si hay un error verificar esto
	col := db.Collection("usuarios")

	u.Password, _ = EncriptarPassword(u.Password) //La password no se guarda tal cual se ingresa

	result, err := col.InsertOne(ctx, u) //El result es un documento JSON que trae varios valores
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil



}