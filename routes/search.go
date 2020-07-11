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

	var id int
	var result []models.Product
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

	id, err = strconv.Atoi(sp.Search)

	/*Si da error es un string e  invocamos un repositorio*/
	if err != nil {

		if len(sp.Search) <= 3 {
			http.Error(w, "El campo de búsquda debe tener al menos 4 caracteres ", 400)
			return
		}

		result, err = repositories.SearchProductByString(sp.Search)

		if err != nil {
			http.Error(w, "Sin resultados "+err.Error(), 400)
			return
		}

	} else {
		//Si no da error es un entero valido y buscamos por su id
		result, err = repositories.SearchProductById(id)

	}

	if err != nil {
		http.Error(w, "Registro no encontrado "+err.Error(), 400)
		return
	}

	if result == nil {
		http.Error(w, "Sin resultados", http.StatusNoContent)
	}

	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)

	fmt.Println(sp.Search)
}
