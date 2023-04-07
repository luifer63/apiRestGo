package main

import (
	"log"
	"github.com/luifer63/apiRestGo/handlers"
	"github.com/luifer63/apiRestGo/bd"
)

func main(){

	if bd.ChequeoConexion() == 0 {
		log.Fatal("Sin conexión a BD")
		return
	}
	handlers.Manejadores()

}