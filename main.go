package main

import (
	"log"

	"github.com/fenriz07/Desafio-W/db"
	"github.com/fenriz07/Desafio-W/handlers"
)

func main() {

	if db.CheckConnection() == 0 {
		log.Fatal("Sin conexión a la bd")
		return
	}

	//Carga los seeders
	//seeder.LoadSeeder()

	handlers.Init()
}
