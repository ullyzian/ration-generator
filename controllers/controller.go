package controllers

import (
	"html/template"
	"net/http"
)

var rootTmpl = template.Must(template.ParseFiles("./templates/base.html", "./templates/index.html"))

func RootHandler(w http.ResponseWriter, r *http.Request) {
	err := rootTmpl.ExecuteTemplate(w, "base", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

var programsTmpl = template.Must(template.ParseFiles("./templates/base.html", "./templates/programs.html"))

func ProgramsHandler(w http.ResponseWriter, r *http.Request) {
	err := programsTmpl.ExecuteTemplate(w, "base", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
