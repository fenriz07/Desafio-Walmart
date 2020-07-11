package seeder

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fenriz07/Desafio-W/db"
	"github.com/fenriz07/Desafio-W/models"
)

/*LoadSeeder Se encarga de cargar el seeder*/
func LoadSeeder() {
	nameFile := "products.json"

	products, err := getContentFile(nameFile)

	if err != nil {
		os.Exit(2)
	}

	for _, product := range products {
		_, _, _ = db.InsertProduct(product)
	}

	fmt.Println("ProductsSeeder Cargado")

}

func getContentFile(file string) ([]models.Product, error) {
	fileRoute := "./seeder/" + file

	binaryJSON, err := ioutil.ReadFile(fileRoute)
	var products []models.Product

	if err != nil {
		return products, err
	}

	err = json.Unmarshal(binaryJSON, &products)

	if err != nil {
		return products, err
	}

	return products, err
}
