package db

import (
	"context"
	"fmt"
	"time"

	"github.com/fenriz07/Desafio-W/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*FindProductID busca un producto por su id*/
func FindProductID(idProduct int) (models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("walmart")
	col := db.Collection("products")
	var product models.Product

	idProductTmp := idProduct

	condition := bson.M{
		"product_id": idProductTmp,
	}

	err := col.FindOne(ctx, condition).Decode(&product)

	if err != nil {
		fmt.Println("Registro no encontrado " + err.Error())
		return product, err
	}

	return product, err
}
