package bd

import (
	"context"
	"time"
	"github.com/Anitori/Project-Go/models"
)

//Graba la relación en la BD 
func InsertoRelacion(t models.Relacion) (bool, error){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("ProjectGo")
	col := db.Collection("relacion") //Colección relacion

	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}

	return true, nil
}