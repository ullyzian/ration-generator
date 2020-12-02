package server

import (
	"github.com/ullyzian/ration-generator/app/models"
	"html/template"
	"net/http"
	"strconv"
)

func (s *Server) handleRoot() http.HandlerFunc {
	tmpl := template.Must(template.ParseFiles("./app/templates/base.html", "./app/templates/index.html"))

	return func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.ExecuteTemplate(w, "base", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) handleGetPrograms() http.HandlerFunc {
	tmpl := template.Must(template.ParseFiles("./app/templates/base.html", "./app/templates/programs.html"))
	type context struct {
		Programs []models.Program
		Err    string
	}

	return func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.ExecuteTemplate(w, "base", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) handleGetDishes() http.HandlerFunc {
	tmpl := template.Must(template.ParseFiles("./app/templates/base.html", "./app/templates/dishes.html"))
	type context struct {
		Dishes []models.Dish
		Err    string
	}

	return func(w http.ResponseWriter, r *http.Request) {
		dishes, err := s.store.Dish().GetAll()
		if err != nil {
			err = tmpl.ExecuteTemplate(w, "base", context{Dishes: nil, Err: err.Error()})
			return
		}
		err = tmpl.ExecuteTemplate(w, "base", &context{Dishes: dishes, Err: ""})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

}

func (s *Server) handleCreateDish() http.HandlerFunc {
	tmpl := template.Must(template.ParseFiles("./app/templates/base.html", "./app/templates/dishesForm.html"))
	type context struct {
		Dish    *models.Dish
		Created bool
		Err     string
	}

	get := func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.ExecuteTemplate(w, "base", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	post := func(w http.ResponseWriter, r *http.Request) {
		title := r.FormValue("title")
		portion, _ := strconv.ParseInt(r.FormValue("portion"), 10, 64)
		calories, _ := strconv.ParseInt(r.FormValue("calories"), 10, 64)
		dish, err := s.store.Dish().Create(&models.Dish{Title: title, Portion: int(portion), Calories: int(calories)})
		if err != nil {
			err = tmpl.ExecuteTemplate(w, "base", &context{Dish: dish, Created: false, Err: err.Error()})
			return
		}
		err = tmpl.ExecuteTemplate(w, "base", &context{Dish: dish, Created: true})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	return func(w http.ResponseWriter, r *http.Request) {
		switch method := r.Method; method {
		case http.MethodGet:
			get(w, r)
		case http.MethodPost:
			post(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}

}

func (s *Server) handleCreateProgram() http.HandlerFunc {
	tmpl := template.Must(template.ParseFiles("./app/templates/base.html", "./app/templates/programsForm.html"))

	type context struct {
		Program *models.Program
		Dishes  []models.Dish
		Created bool
		Err     string
	}

	get := func(w http.ResponseWriter, r *http.Request, dishes *[]models.Dish) {
		if err := tmpl.ExecuteTemplate(w, "base", context{Dishes: *dishes}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	post := func(w http.ResponseWriter, r *http.Request, dishes *[]models.Dish ) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		name := r.FormValue("name")
		dishesIds := r.Form["dishes"]
		dishesByIds, err := s.store.Dish().GetByIds(dishesIds)
		if err != nil {
			err = tmpl.ExecuteTemplate(w, "base", context{Created: false, Err: err.Error(), Dishes: *dishes})
			return
		}
		program, err := s.store.Program().Create(&models.Program{Name: name, Dishes: dishesByIds})
		if err != nil {
			err = tmpl.ExecuteTemplate(w, "base", context{Created: false, Err: err.Error(), Dishes: *dishes})
			return
		}
		if err := tmpl.ExecuteTemplate(w, "base", &context{Created: true, Program: program, Dishes: *dishes}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	return func(w http.ResponseWriter, r *http.Request) {
		dishes, err := s.store.Dish().GetAll()
		if err != nil {
			err = tmpl.ExecuteTemplate(w, "base", context{Err: err.Error()})
			return
		}
		switch method := r.Method; method {
		case http.MethodGet:
			get(w, r, &dishes)
		case http.MethodPost:
			post(w, r, &dishes)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
