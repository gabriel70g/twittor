package routers

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gabriel70g/twittor/bd"
)

/*ObtenerBanner obtener avatar*/
func ObtenerBanner(w http.ResponseWriter, r *http.Request) {
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

	var mipath, _ = os.Getwd()

	openFile, err := os.Open(mipath + "/uploads/banners/" + perfil.Banner)
	if err != nil {
		fmt.Println(err)
		fmt.Println(mipath + "/uploads/banners/")
		http.Error(w, "Imagen no encontrada"+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, openFile)
	if err != nil {
		http.Error(w, "Error al copiar la imagen", http.StatusBadRequest)
	}
}
