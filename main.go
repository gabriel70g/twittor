package main

import (
	"log"

	"github.com/gabriel70g/twittor/bd"
	"github.com/gabriel70g/twittor/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}
