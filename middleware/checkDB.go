package middleware

import (
	"net/http"

	"github.com/fenriz07/Desafio-W/db"
)

/*CheckDB Middleware que chequea la conexión a la bd*/
func CheckDB(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == 0 {
			http.Error(w, "Conexión perdida con la base de datos", 500)
			return
		}

		next.ServeHTTP(w, r)
	}
}
