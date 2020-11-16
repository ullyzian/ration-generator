package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main()  {
	fs := http.FileServer(http.Dir("app/static"))

	http.HandleFunc("/", rootHandler)
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/programs", programsHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func render(w http.ResponseWriter, templateName string)  {
	tmpl, err := template.ParseFiles("app/templates/base.html", fmt.Sprintf("app/templates/%v.html", templateName))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.ExecuteTemplate(w, "base", "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "index")
}

func programsHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "programs")
}