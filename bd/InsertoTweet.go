package bd

import (
	"context"
	"time"

	"github.com/luifer63/apiRestGo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
// InsertoTweet -> Graba tweet en la BD
func InsertoTweet(t models.GraboTweet)(string, bool, error){
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("apirestgoremote")
	col := db.Collection("tweets")

	registro := bson.M{
		"userid": t.UserId,
		"mensaje": t.Mensaje,
		"fecha": t.Fecha,
	}

	result, err := col.InsertOne(ctx, registro)

	if err!= nil {
		return "", false, err
	}

	objID , _ := result.InsertedID.(primitive.ObjectID)

	return objID.String(), true, nil
}