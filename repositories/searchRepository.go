package repositories

import (
	"github.com/fenriz07/Desafio-W/db"
	"github.com/fenriz07/Desafio-W/models"
)

/*SearchProductById busca un producto por su id*/
func SearchProductById(idProduct int) ([]models.Product, error) {
	return db.FindProductID(idProduct)
}

/*SearchProductByString busca un producto por un texto enviado*/
func SearchProductByString(text string) ([]models.Product, error) {
	return db.FindProductString(text, 1)
}
