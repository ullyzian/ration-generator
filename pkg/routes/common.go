package routes

import (
	"html/template"
	"net/http"
	"github.com/ullyzian/ration-generator/pkg/server"
)

func handleRoot(s *server.Server) http.HandlerFunc {
	tmpl := template.Must(template.ParseFiles("./pkg/templates/base.html", "./pkg/templates/index.html"))

	return func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.ExecuteTemplate(w, "base", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}