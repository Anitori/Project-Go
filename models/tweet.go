package models

//Captura del body, el mensaje que nos llega
type Tweet struct {
	Mensaje string `bson:"mensaje" json:"mensaje"`
}
