package db

import (
	"context"
	"fmt"
	"time"

	"github.com/fenriz07/Desafio-W/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*FindProductString busca un producto por su id*/
func FindProductString(text string, pagina int64) ([]models.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("walmart")
	col := db.Collection("products")

	var products []models.Product

	textTmp := text

	/*condition := bson.M{
		"brand":       bson.M{"$regex": `(?i)` + textTmp},
		"description": bson.M{"$regex": `(?i)` + textTmp},
	}*/

	condition := bson.M{
		"$or": []interface{}{
			bson.M{"brand": bson.M{"$regex": `(?i)` + textTmp}},
			bson.M{"description": bson.M{"$regex": `(?i)` + textTmp}},
		},
	}

	options := options.Find()
	options.SetLimit(20)
	options.SetSort(bson.D{{Key: "product_id", Value: 1}})
	options.SetSkip((pagina - 1) * 20)

	cur, err := col.Find(ctx, condition, options)

	if err != nil {
		fmt.Println("Registro no encontrado " + err.Error())
		return products, err
	}

	for cur.Next(ctx) {
		var p models.Product
		err := cur.Decode(&p)

		if err != nil {
			fmt.Println("epa_4")
			//fmt.Println(err.Error())
			return products, err
		}

		if p.IsProductPalindrome() {
			p.ApplyDiscount()
		}

		products = append(products, p)
	}

	return products, err
}
