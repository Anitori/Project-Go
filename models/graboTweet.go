package models 

import "time"

//Formato o estructura que tendrá nuestro tweet
type GraboTweet struct{
	UserID string `bson:"userid" json:"userid,omitempty"`
	Mensaje string `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha time.Time `bson:"fecha" json:"fecha,omitempty"`
}