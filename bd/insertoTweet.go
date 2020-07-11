package bd

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertoTweet crea un nuevo tweet*/
func InsertoTweet() (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("walmart")
	col := db.Collection("mensajes")

	registro := bson.M{
		"texto": "Probando",
	}

	result, err := col.InsertOne(ctx, registro)

	if err != nil {
		return string(""), false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)

	return objID.String(), true, nil
}
