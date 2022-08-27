package bd

import (
	"context"
	"fmt"
	"time"
	"github.com/Anitori/Project-Go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)


//Esta función va a devolver un slice y un booleano
func LeoUsuariosTodos(ID string, page int64, search string, tipo string) ([]*models.Usuario, bool){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("ProjectGo")
	col := db.Collection("usuarios") 


	var results []*models.Usuario //Voy a enviar al http un slice de usuario


	//Seteo las opciones de busqueda
	findOptions := options.Find()
	findOptions.SetSkip((page -1)*20)
	findOptions.SetLimit(20) //Voy a devolver solo 20 resultados


	//Seteo la condición
	query := bson.M{
		"nombre": bson.M{"$regex": `(?i)`+ search},  //la i se fija en que no importa si es mayuscula o minuscula
	}


	//Realizo la busqueda y me devuelve un cursor
	cur, err := col.Find(ctx, query, findOptions)
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	var encontrado, incluir bool
	
	for cur.Next(ctx){
		var s models.Usuario 
		err := cur.Decode(&s)
		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}


		var r models.Relacion
		r.UsuarioID=ID
		r.UsuarioRelacionID=s.ID.Hex()

		incluir = false 

		encontrado, err = ConsultoRelacion(r)
		if tipo == "new" && encontrado == false {
			incluir=true
		}

		if tipo == "follow" && encontrado == false {
			incluir=true
		}

		if r.UsuarioRelacionID == ID {
			incluir = false
		}

		if incluir == true{
			s.Password=""
			s.Biografia=""
			s.SitioWeb=""
			s.Ubicacion=""
			s.Banner=""
			s.Email=""

			results = append(results, &s)
		}

	}


	err = cur.Err()
	if err != nil{
		fmt.Println(err.Error())
		return results, false
	}

	cur.Close(ctx)
	return results, true
	
}