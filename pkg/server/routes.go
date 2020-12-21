package server

import (
	"github.com/gorilla/mux"
	"github.com/ullyzian/ration-generator/pkg/models"
	"html/template"
	"net/http"
	"strconv"
)

func (s *Server) handleGenerator() http.HandlerFunc {
	tmpl := template.Must(template.ParseFiles("./pkg/templates/base.html", "./pkg/templates/generator.html"))

	get := func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.ExecuteTemplate(w, "base", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	post := func(w http.ResponseWriter, r *http.Request) {
		//purposeF := r.FormValue("purpose")
		//contraindicationsF := r.FormValue("contradications")
		//fmt.Println(purposeF, contraindicationsF)
		//dishes, err := s.store.Dish().GetAll()
		//if err != nil {
		//	http.Error(w, err.Error(), http.StatusInternalServerError)
		//}
		//var filteredDishes []models.Dish
		//for i := range dishes {
		//	if dishes[i].
		//}

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

func (s *Server) handleProgramsList() http.HandlerFunc {
	tmpl := template.Must(template.ParseFiles("./pkg/templates/base.html", "./pkg/templates/programs.html"))
	type context struct {
		Programs []models.Program
		Err      string
	}

	return func(w http.ResponseWriter, r *http.Request) {
		programs, err := s.store.Program().GetAll()
		if err != nil {
			err = tmpl.ExecuteTemplate(w, "base", context{Programs: nil, Err: err.Error()})
		}
		err = tmpl.ExecuteTemplate(w, "base", context{Programs: programs, Err: ""})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (s *Server) handleRoot() http.HandlerFunc {
	tmpl := template.Must(template.ParseFiles("./pkg/templates/base.html", "./pkg/templates/index.html"))

	return func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.ExecuteTemplate(w, "base", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (s *Server) handleGetDishes() http.HandlerFunc {
	tmpl := template.Must(template.ParseFiles("./pkg/templates/base.html", "./pkg/templates/dishes/list.html"))
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
	tmpl := template.Must(template.ParseFiles("./pkg/templates/base.html", "./pkg/templates/dishes/create.html"))
	type context struct {
		Dish *models.Dish
		Err  string
	}

	get := func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.ExecuteTemplate(w, "base", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	post := func(w http.ResponseWriter, r *http.Request) {
		title := r.FormValue("title")
		contr := r.FormValue("contradictions")
		portion, _ := strconv.ParseInt(r.FormValue("portion"), 10, 64)
		calories, _ := strconv.ParseInt(r.FormValue("calories"), 10, 64)
		dish, err := s.store.Dish().Create(&models.Dish{Title: title, Portion: int(portion), Calories: int(calories), Contraindications: contr})
		if err != nil {
			err = tmpl.ExecuteTemplate(w, "base", &context{Dish: dish, Err: err.Error()})
			return
		}
		http.Redirect(w, r, "http://localhost:8080/dishes", 301)
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

func (s *Server) handleEditDish() http.HandlerFunc {
	tmpl := template.Must(template.ParseFiles("./pkg/templates/base.html", "./pkg/templates/dishes/edit.html"))
	type context struct {
		Dish *models.Dish
		Err  string
	}

	get := func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])
		dish, err := s.store.Dish().GetById(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if err = tmpl.ExecuteTemplate(w, "base", context{Dish: dish}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	post := func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])
		title := r.FormValue("title")
		contr := r.FormValue("contradictions")
		portion, _ := strconv.Atoi(r.FormValue("portion"))
		calories, _ := strconv.Atoi(r.FormValue("calories"))
		d := &models.Dish{Id: id, Title: title, Portion: int(portion), Calories: int(calories), Contraindications: contr}
		_, err := s.store.Dish().Edit(d)
		if err != nil {
			err = tmpl.ExecuteTemplate(w, "base", &context{Dish: d, Err: err.Error()})
		}
		http.Redirect(w, r, "/dishes", http.StatusSeeOther)
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

func (s *Server) handleDeleteDish() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])
		if _, err := s.store.Dish().Delete(id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		http.Redirect(w, r, "/dishes", http.StatusSeeOther)
	}
}

func (s *Server) handleGetPrograms() http.HandlerFunc {
	tmpl := template.Must(template.ParseFiles("./pkg/templates/base.html", "./pkg/templates/programs/list.html"))
	type context struct {
		Programs []models.Program
		Err      string
	}

	return func(w http.ResponseWriter, r *http.Request) {
		programs, err := s.store.Program().GetAll()
		if err != nil {
			err = tmpl.ExecuteTemplate(w, "base", context{Programs: nil, Err: err.Error()})
		}
		err = tmpl.ExecuteTemplate(w, "base", context{Programs: programs, Err: ""})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (s *Server) handleCreateProgram() http.HandlerFunc {
	tmpl := template.Must(template.ParseFiles("./pkg/templates/base.html", "./pkg/templates/programs/create.html"))

	type context struct {
		Program *models.Program
		Dishes  []models.Dish
		Err     string
	}

	get := func(w http.ResponseWriter, r *http.Request, dishes *[]models.Dish) {
		if err := tmpl.ExecuteTemplate(w, "base", context{Dishes: *dishes}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	post := func(w http.ResponseWriter, r *http.Request, dishes *[]models.Dish) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		name := r.FormValue("name")
		dishesIds := r.Form["dishes"]
		dishesByIds, err := s.store.Dish().GetByIds(dishesIds)
		if err != nil {
			err = tmpl.ExecuteTemplate(w, "base", context{Err: err.Error(), Dishes: *dishes})
			return
		}
		_, err = s.store.Program().Create(&models.Program{Name: name, Dishes: dishesByIds})
		if err != nil {
			err = tmpl.ExecuteTemplate(w, "base", context{Err: err.Error(), Dishes: *dishes})
			return
		}
		http.Redirect(w, r, "/programs", http.StatusSeeOther)
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

func (s *Server) handleEditProgram() http.HandlerFunc {
	tmpl := template.Must(template.ParseFiles("./pkg/templates/base.html", "./pkg/templates/programs/edit.html"))
	type context struct {
		Program *models.Program
		Dishes  []models.Dish
		Err     string
	}

	get := func(w http.ResponseWriter, r *http.Request, dishes *[]models.Dish) {
		if err := tmpl.ExecuteTemplate(w, "base", context{Dishes: *dishes}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	post := func(w http.ResponseWriter, r *http.Request, dishes *[]models.Dish) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		name := r.FormValue("name")
		dishesIds := r.Form["dishes"]
		dishesByIds, err := s.store.Dish().GetByIds(dishesIds)
		if err != nil {
			err = tmpl.ExecuteTemplate(w, "base", context{Err: err.Error(), Dishes: *dishes})
			return
		}
		_, err = s.store.Program().Create(&models.Program{Name: name, Dishes: dishesByIds})
		if err != nil {
			err = tmpl.ExecuteTemplate(w, "base", context{Err: err.Error(), Dishes: *dishes})
			return
		}
		http.Redirect(w, r, "/programs", http.StatusSeeOther)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		dishes, err := s.store.Dish().GetAll()
		if err != nil {
			err = tmpl.ExecuteTemplate(w, "base", context{Err: err.Error()})
			return
		}
		//vars := mux.Vars(r)
		//id, _ := strconv.Atoi(vars["id"])
		//program, err := s.store.Program().GetBy
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

func (s *Server) handleDeleteProgram() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])
		if _, err := s.store.Program().Delete(id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		http.Redirect(w, r, "/programs", http.StatusSeeOther)
	}
}
