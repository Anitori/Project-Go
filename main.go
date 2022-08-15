package main 


import(
	"log"
	"github.com/Anitori/Project-Go/handlers"
	"github.com/Anitori/Project-Go/bd"
)

func main (){

	if bd.ChequeoConecction() == 0 {
		log.Fatal ("Sin conexión a la BD")
		return
	}

	handlers.Manejadores()
}
