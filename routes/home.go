package routes

import (
	"net/http"
	"text/template"
)

/*Home controlador que renderiza el html*/
func Home(w http.ResponseWriter, r *http.Request) {

	views := template.Must(template.ParseGlob("resources/views/*"))

	err := views.ExecuteTemplate(w, "home.html", nil)

	if err != nil {
		http.Error(w, err.Error(), 500)
	}

}
