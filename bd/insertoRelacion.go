package bd

import (
	"context"
	"time"
	"github.com/luifer63/apiRestGo/models"
)

// InsertoRelacion -> Funcion para insertar una relacion entre usuarios
func InsertoRelacion(t models.Relacion)(bool, error){

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("apirestgoremote")
	col := db.Collection("relaciones")

	_, err := col.InsertOne(ctx, t)

	if err != nil {
		return false, err
	}
	return true, nil
}