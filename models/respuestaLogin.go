package models

/*RepuestaLogin tiene el token que devuelve con el log*/
type RespuestaLogin struct {
	Token string `json:"token, omitempty"`
}
