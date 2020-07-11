package db

import (
	"context"
	"fmt"
	"time"

	"github.com/fenriz07/Desafio-W/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*FindProductID busca un producto por su id*/
func FindProductID(idProduct int) ([]models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("walmart")
	col := db.Collection("products")

	var products []models.Product

	idProductTmp := idProduct

	condition := bson.M{
		"product_id": idProductTmp,
	}

	cur, err := col.Find(ctx, condition)

	if err != nil {
		fmt.Println("Registro no encontrado " + err.Error())
		return products, err
	}

	for cur.Next(ctx) {
		var p models.Product
		err := cur.Decode(&p)

		if err != nil {
			fmt.Println(err.Error())
			return products, err
		}

		if p.IsProductPalindrome() {
			p.ApplyDiscount()
		}

		products = append(products, p)
	}

	return products, err
}
