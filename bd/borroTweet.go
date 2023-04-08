package bd

import (
	"context"
	"time"
	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"

)

// BorroTweet -> Borra un tweet determinado
func BorroTweet(ID string, UserID string) error{

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 15)
	defer cancel()

	db := MongoCN.Database("apirestgoremote")
	col := db.Collection("tweets")

	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id": objID,
		"userid": UserID,
	}
	_, err := col.DeleteOne(ctx, condicion)
	return err
}