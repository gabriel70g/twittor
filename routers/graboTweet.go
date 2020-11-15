package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gabriel70g/twittor/bd"
	"github.com/gabriel70g/twittor/models"
)

/*GraboTweet Grabo los Tweets */
func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	var status bool
	_, status, err = bd.InsertoTweet(registro)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar insertar el registro, reintente nuevamente"+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado isertar el Tweet", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
