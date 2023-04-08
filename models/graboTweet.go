package models

import "time"

// GraboTweet -> Formato que tendrá el tweet en BD
type GraboTweet struct {
	UserId string `bson:"userid" json:"userid,omitempty"`
	Mensaje string `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha time.Time `bson:"fecha" json:"fecha,omitempty"`
}

