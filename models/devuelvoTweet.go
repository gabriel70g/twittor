package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*DevuelvoTweets es la estructura con la que devolvemos los Tweets */
type DevuelvoTweets struct {
	ID      primitive.ObjectID `bson:"_id" json:"_id,omitwmpty"`
	UserID  string             `bson:"userid" json:"userId,omitwmpty"`
	Mensaje string             `bson:"mensaje" json:"mensaje,omitwmpty"`
	Fecha   time.Time          `bson:"fecha" json:"fecha,omitwmpty"`
}
