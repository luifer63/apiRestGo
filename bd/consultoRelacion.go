package bd

import (
	"context"
	"time"
	"fmt"
	"github.com/luifer63/apiRestGo/models"
	"go.mongodb.org/mongo-driver/bson"
)
//ConsultoRelacion -> Saber que relación tienen dos usuarios
func ConsultoRelacion(t models.Relacion)(bool, error){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("apirestgoremote")
	col := db.Collection("relaciones")

	condicion := bson.M{
		"usuarioid": t.UsuarioID,
		"usuariorelacionid": t.UsuarioRelacionID,
	}

	var resultado models.Relacion

	fmt.Println(resultado)
	err := col.FindOne(ctx, condicion).Decode(&resultado)

	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	return true, nil
}