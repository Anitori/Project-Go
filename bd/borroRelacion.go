package bd

import (
	"context"
	"time"
	"github.com/Anitori/Project-Go/models"
)

//Borra la relación en la BD
func BorroRelacion(t models.Relacion) (bool, error){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("ProjectGo")
	col := db.Collection("relacion") 

	_, err := col.DeleteOne(ctx,t)
	if err != nil {
		return false, err
	}

	return true, nil
}