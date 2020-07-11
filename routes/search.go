package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/fenriz07/Desafio-W/models"
	"github.com/fenriz07/Desafio-W/repositories"
)

/*SearchProducts controlador de la ruta search*/
func SearchProducts(w http.ResponseWriter, r *http.Request) {

	//w.Header().Add("content-type", "application/json")

	var sp models.SearchRequest

	err := json.NewDecoder(r.Body).Decode(&sp)

	if err != nil {
		http.Error(w, "Error en la búsqueda "+err.Error(), 400)
		return
	}

	if len(sp.Search) == 0 {
		http.Error(w, "Favor pase un parametro a buscar", 400)
		return
	}

	var id int

	id, err = strconv.Atoi(sp.Search)

	/*Si da error es un string e  invocamos un repositorio*/
	if err != nil {
		fmt.Println(err.Error())
	}

	/*Si es un entero entonces invocamos un repositorio para buscar el producto por su id*/
	result, err := repositories.SearchProductById(id)

	if err != nil {
		http.Error(w, "Registro no encontrado "+err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)

	fmt.Println(sp.Search)
}
