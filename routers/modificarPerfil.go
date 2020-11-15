package routers

import (
	"encoding/json"
	"net/http"

	"github.com/gabriel70g/twittor/bd"
	"github.com/gabriel70g/twittor/models"
)

/*ModificarPerfil modifica el perfil de usuario*/
func ModificarPerfil(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos incorrectos: "+err.Error(), http.StatusBadRequest)
		return
	}
	var status bool
	status, err = bd.ModificoRegistro(t, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar modificar el registro. Reintente nuevamente:"+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado modifica el registro del usuario:", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
