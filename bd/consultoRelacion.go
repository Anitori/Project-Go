package bd

import (
	"context"
	"time"
	"fmt"

	"github.com/Anitori/Project-Go/models"
	"go.mongodb.org/mongo-driver/bson"
)

//Consulta la relación entre dos usuarios

func ConsultoRelacion(t models.Relacion) (bool, error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("ProjectGo")
	col := db.Collection("relacion") 

	condicion := bson.M{
		"usuarioid": t.UsuarioID,
		"usuariorelacionid": t.UsuarioRelacionID,
	}

	var resultado models.Relacion

	fmt.Println(resultado)
	err := col.FindOne(ctx, condicion).Decode(&resultado)

	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	return true, nil

}