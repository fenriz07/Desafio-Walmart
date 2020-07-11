package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/fenriz07/Desafio-W/middleware"
	"github.com/fenriz07/Desafio-W/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Init punto de entrada que define las diferentes rutas del api*/
func Init() {
	router := mux.NewRouter()

	router.HandleFunc("/search", middleware.CheckDB(routes.SearchProducts)).Methods("POST")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
