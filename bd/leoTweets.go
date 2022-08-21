package bd 

import (
	"context"
	"time"
	"log"
	"github.com/Anitori/Project-Go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func LeoTweets(ID string, pagina int64) ([]*models.DevuelvoTweets, bool) {  

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("ProjectGo")
	col := db.Collection("tweet") //Colección tweet

	var resultados []*models.DevuelvoTweets


	//Voy a leer mi tabla tweet donde mi userID sea el que me llegó por parametro
	condicion := bson.M{
		"userid": ID,
	}
	
	opciones := options.Find()
	opciones.SetLimit(20) //Cuantos documentos me tiene que traer límite. Como traer de a 20 mensajes por página
	opciones.SetSort(bson.D{{Key:"fecha", Value: -1}}) //Como se van a ordenar, el -1 es en descendente (Los últimos primero)
	opciones.SetSkip((pagina -1)*20)

	//Un cursor es un puntero, es como si fuera una tabla de BD, donde se van a grabar los resultados y los voy a poder recorrer de a uno.

	cursor, err := col.Find(ctx, condicion, opciones)
	if err != nil {
		log.Fatal(err.Error())
		return resultados, false
	}

	for cursor.Next(context.TODO()){
		
		var registro models.DevuelvoTweets
		err := cursor.Decode(&registro)
		if err != nil {
			return resultados, false
		}

		resultados = append(resultados, &registro)
	}

	return resultados, true


}

