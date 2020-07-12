package db

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN tiene el cliente de la conexión*/
var MongoCN = conection()

var dbHost = os.Getenv("DB_HOST")
var dbPort = os.Getenv("DB_PORT")

var clientOptions = options.Client().ApplyURI("mongodb://" + dbHost + ":" + dbPort)

//var clientOptions = options.Client().ApplyURI("mongodb+srv://fenriz:ReGBipJ9yLE3hs3@cluster0-42e6i.mongodb.net/twittor?retryWrites=true&w=majority")

/* conection es la funcion que hace la conexion*/
func conection() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Conexión exitosa con la bd")
	return client

}

/*CheckConnection Comprueba la conexión*/
func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)

	if err != nil {
		return 0
	}

	return 1
}
