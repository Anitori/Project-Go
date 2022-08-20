package bd 

import (
	"context"
	"fmt"
	"time"
	"github.com/Anitori/Project-Go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Busca un perfil en la BD
func BuscoPerfil(ID string) (models.Usuario, error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("ProjectGo")
	col := db.Collection("usuarios")

	var perfil models.Usuario
	objID, _ := primitive.ObjectIDFromHex(ID)


	condicion := bson.M{
		"_id": objID,
	}

	err:=col.FindOne(ctx, condicion).Decode(&perfil)  //Devolución de la información
	perfil.Password=""
	if err != nil {
		fmt.Println("Registro no encontrado "+err.Error())
		return perfil, err
	}
	return perfil, nil 
} 