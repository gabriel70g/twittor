package routers

import (
	"net/http"

	"github.com/gabriel70g/twittor/bd"
	"github.com/gabriel70g/twittor/models"
)

/*BajaRalacion dar de baja la relacion */
func BajaRalacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.BorroRelacion(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar borrar realción", http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado borrar la realción", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
