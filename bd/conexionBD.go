package bd

import (
	"context"
	"log" //Grabar nombre de texto dentro del log de ejecución
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


var MongoCN = ConectarBD() //MongoCN es el objeto de conexión a la BD 
var clientOptions = options.Client().ApplyURI("mongodb+srv://elexicon:polifemo131@cluster0.skkxwd0.mongodb.net/?retryWrites=true&w=majority") //Setear la URL de la base de datos 


func ConectarBD() *mongo.Client { //Esta es la función que me permite conectar la BD
	client, err := mongo.Connect(context.TODO(), clientOptions) //Hace una conexión a la base de datos,toma la conexión de clientOptions 
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexión Exitosa con la BD")
	return client
}

func ChequeoConecction() int {
	err := MongoCN.Ping(context.TODO(), nil) // En todas las demas rutinas cuando yo tenga que hablar de la conexión, tengo que hablar de MongoCN
											// voy a poder hacer todas las operaciones de bases de datos con MongoCN
	if err != nil{
		return 0
	}
	return 1
}
