package bd

import (
	"context"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func BorroTweet(ID string, UserID string) error { //el ID es el del tweet, y el otro del usuario
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("ProjectGo")
	col := db.Collection("tweet") //Colección tweet

	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		":id": objID,
		"userid": UserID,
	}

	_, err := col.DeleteOne(ctx, condicion)
	return err 
} 