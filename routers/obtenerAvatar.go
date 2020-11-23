package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gabriel70g/twittor/bd"
)

/*ObtenerAvatar obtener avatar*/
func ObtenerAvatar(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro id", http.StatusBadRequest)
		return
	}
	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "Usuario no encontrado", http.StatusBadRequest)
		return
	}
	if len(strings.TrimSpace(perfil.Avatar)) == 0 {
		http.Error(w, "El usuario no tiene avatar", http.StatusBadRequest)
		return
	}
	var mipath, _ = os.Getwd()

	openFile, err := os.Open(mipath + "/uploads/avatars/" + perfil.Avatar)
	if err != nil {
		http.Error(w, "Imagen no encontrada", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, openFile)
	if err != nil {
		http.Error(w, "Error al copiar la imagen"+err.Error(), http.StatusBadRequest)
	}
}
