package routers

import (
	"net/http"

	"github.com/gabriel70g/twittor/bd"
	"github.com/gabriel70g/twittor/models"
)

/*AltaRelacion realizo el alta de una relacion*/
func AltaRelacion(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parametro ID es obligatorio", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.InsertoRelacion(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar isertar realción", http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado isertar la realción", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
