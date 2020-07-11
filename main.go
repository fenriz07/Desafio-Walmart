package main

import (
	"log"

	"github.com/fenriz07/Desafio-W/bd"
	"github.com/fenriz07/Desafio-W/handlers"
)

func main() {

	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la bd")
		return
	}

	handlers.Init()
}
