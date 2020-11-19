package models

/*RespuestaConsultaRelacion tiene el true o el false que
se obtiene de cosultar la relacion de 2 usuairos*/
type RespuestaConsultaRelacion struct {
	Status bool `json:"status"`
}
