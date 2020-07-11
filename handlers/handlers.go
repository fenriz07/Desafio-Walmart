package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/fenriz07/Desafio-W/bd"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Init punto de entrada que define las diferentes rutas del api*/
func Init() {
	router := mux.NewRouter()
	router.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusCreated)

		type Message struct {
			Name string
			Body string
		}

		mensaje, estado, err := bd.InsertoTweet()

		var m = Message{}

		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		if !estado {
			http.Error(w, err.Error(), 400)
			return
		} else {
			m = Message{"Probando", mensaje}
		}

		json.NewEncoder(w).Encode(m)

	}).Methods("GET")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
