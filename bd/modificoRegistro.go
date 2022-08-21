package bd

import (
	"context"
	"time"

	"github.com/Anitori/Project-Go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Permite modificar el perfil del usuario
func ModificoRegistro(u models.Usuario, ID string) (bool, error){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("ProjectGo")
	col := db.Collection("usuarios")


	//Creo una interfaz vacía, y le digo que cuando estoy creando una clave dentro de ese mapa para el campo nombre le grabo un valor

	//Creo un registro vacío, que es un mapa de tipo interface, luego le gravo las clave valor que haga falta. Despues lo paso de parámetro en la instrucción de mongo $set
	registro := make(map[string]interface{})
	if len(u.Nombre) > 0{
		registro["nombre"] = u.Nombre
	}
	if len(u.Apellidos) > 0{
		registro["apellidos"] = u.Apellidos
	}

	registro["fechaNacimiento"] = u.FechaNacimiento

	if len(u.Avatar) > 0{
		registro["avatar"] = u.Avatar
	}
	if len(u.Banner) > 0{
		registro["banner"] = u.Banner
	}
	if len(u.Biografia) > 0{
		registro["biografia"] = u.Biografia
	}
	if len(u.Ubicacion) > 0{
		registro["ubicacion"] = u.Ubicacion
	}
	if len(u.SitioWeb) > 0{
		registro["sitioWeb"] = u.SitioWeb
	}



	//Seteo de los campos
	updtString := bson.M{
		"$set": registro, 
	}


	//Esto va a convertir mi ID que es un string en un formato OBJECT ID
	objID, _ := primitive.ObjectIDFromHex(ID)

	//Agregar un filtro
	filtro :=bson.M{"_id": bson.M{"$eq":objID}}

	_, err := col.UpdateOne(ctx, filtro, updtString)
	if err != nil{
		return false, err
	}

	return true, nil

}