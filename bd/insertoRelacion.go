package bd

import (
	"context"
	"time"

	"github.com/gabriel70g/twittor/models"
)

/*InsertoRelacion inserto una relacion en la tabla*/
func InsertoRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("relation")

	_, err := col.InserOne(ctx, t)
	if err != nil {
		return false, err
	}

	return true, nil
}
