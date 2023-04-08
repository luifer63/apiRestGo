package bd

import (
	"context"
	"time"
	"github.com/luifer63/apiRestGo/models"
)

//BorroRelacion -> funcion para dar de baja una relacion de usuarios
func BorroRelacion(t models.Relacion)(bool, error){

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("apirestgoremote")
	col := db.Collection("relaciones")

	_, err := col.DeleteOne(ctx, t)

	if err != nil {
		return false, err
	}
	return true, nil
}