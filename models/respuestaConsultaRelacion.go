package models 


//Tiene el true o false que se obtiene de consultar la relación entre 2 usuarios
type RespuestaConsultaRelacion struct {
	Status bool `json:"status"`
}