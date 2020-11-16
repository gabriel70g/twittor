package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/gabriel70g/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*BuscoPerfil Busco perril*/
func BuscoPerfil(ID string) (models.Usuario, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	db := MongoCN.Database("twittor")
	col := db.Collection("users")

	defer cancel()

	var perfil models.Usuario
	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condicion).Decode(&perfil)
	perfil.Password = ":)"

	if err != nil {
		fmt.Println("Registro no encontrado: " + err.Error())
		return perfil, err
	}

	return perfil, nil
}
