package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Usuario es el modelo de usuario de la base de mongo
type Usuario struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre          string             `bson:"nombre" json:"nombre"`
	Apellidos       string             `bson:"apellidos" json:"apellidos"`
	FechaNacimiento time.Time          `bson:"fechaNacimiento" json:"fechaNacimiento"`
	Email           string             `bson:"email" json:"email"`
	Password        string             `bson:"password" json:"password,omitempty"`
	Avatar          string             `bson:"avatar" json:"avatar,omitempty"`
	Banner          string             `bson:"banner" json:"banner,omitempty"`
	Biografia       string             `bson:"biografia" json:"biografia,omitempty"`
	SitioWeb        string             `bson:"sitiWeb" json:"sitiWeb,omitempty"`
}
