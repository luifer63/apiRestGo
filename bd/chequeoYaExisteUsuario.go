package bd

import (
	"context"
	"time"
	"github.com/luifer63/apiRestGo/models"
	"go.mongodb.org/mongo-driver/bson"
)
// ChequeoYaExisteUsuario recibe el email para ver si el usuario ya está registrado
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("apirestgoremote")
	col := db.Collection("usuarios")

	condicion := bson.M{"email":email}
	var resultado models.Usuario

	err := col.FindOne(ctx, condicion).Decode(&resultado)

	ID := resultado.ID.Hex()

	if err != nil{
		return resultado, false, ID
	}
	return resultado, true, ID
}