package bd 

import (
	"context"
	"time"

	"github.com/Anitori/Project-Go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Graba el tweet en la BD
func InsertoTweet(t models.GraboTweet) (string, bool, error){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("ProjectGo")
	col := db.Collection("tweet") //Colección tweet


	registro := bson.M{
		"userid": t.UserID,
		"mensaje": t.Mensaje,
		"fecha": t.Fecha,
	}

	result, err := col.InsertOne(ctx, registro)
	if err != nil {
		return "", false, err 
	}


	objID, _ := result.InsertedID.(primitive.ObjectID) //La función de primitive devuelve un objebID y un parámetro que no me interesa ahora
	return objID.String(), true, nil
}