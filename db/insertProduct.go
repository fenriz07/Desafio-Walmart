package db

import (
	"context"
	"time"

	"github.com/fenriz07/Desafio-W/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertProduct inserta una nueva fila en la colección products*/
func InsertProduct(products models.Product) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("walmart")
	col := db.Collection("products")

	result, err := col.InsertOne(ctx, products)

	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
