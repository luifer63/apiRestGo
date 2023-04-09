package bd

import (
	"context"
	"time"

	"github.com/luifer63/apiRestGo/models"
	"go.mongodb.org/mongo-driver/bson"

)

//LeoTweetsSeguidores -> funcion para listar los tweets de los seguidores
func LeoTweetsSeguidores(ID string,  page int)([]*models.DevuelvoTweetsSeguidores, bool){

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("apirestgoremote")
	col := db.Collection("relaciones")

	skip := (page -1) * 20

	condiciones := make([]bson.M, 0)

	condiciones = append(condiciones, bson.M{"$match": bson.M{"usuarioid": ID}})
	condiciones = append(condiciones, bson.M{
		"$lookup": bson.M{
			"from": "tweets",
			"localField": "usuariorelacionid",
			"foreignField": "userid",
			"as": "tweet",

		},
	})

	condiciones = append(condiciones, bson.M{"$unwind": "$tweet"})
	condiciones = append(condiciones, bson.M{"$sort":bson.M{"fecha": -1}})
	condiciones = append(condiciones, bson.M{"$skip": skip})
	condiciones = append(condiciones, bson.M{"$limit": 20})

	cursor, _ := col.Aggregate(ctx, condiciones)

	var result []*models.DevuelvoTweetsSeguidores

	err := cursor.All(ctx, &result)

	if err != nil{
		return result, false
	}

	return result, true
}