package controllers

import (
	"html/template"
	"net/http"
	"ration-generator/db"
	"ration-generator/models"
	"strconv"
)

var createDishTmpl = template.Must(template.ParseFiles("./templates/base.html", "./templates/dishesForm.html"))

type CreationTemplateData struct {
	Dish *models.Dish
	Created bool
	Err string
}

func CreateDish	(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		err := createDishTmpl.ExecuteTemplate(w, "base", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else if r.Method == http.MethodPost {
		conn := db.Connect()
		title := r.FormValue("title")
		portion, _ := strconv.ParseInt(r.FormValue("portion"), 10, 64)
		calories, _ :=  strconv.ParseInt(r.FormValue("calories"), 10, 64)
		dish, err := models.CreateDish(conn, title, int(portion), int(calories))
		if err != nil {
			err = createDishTmpl.ExecuteTemplate(w, "base", &CreationTemplateData{Dish: dish, Created: false, Err: err.Error()})
			return
		}
		err = createDishTmpl.ExecuteTemplate(w, "base", &CreationTemplateData{Dish: dish, Created: true})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

var getDishesTmpl = template.Must(template.ParseFiles("./templates/base.html", "./templates/dishes.html"))

type DishesTemplateData struct {
	Dishes []models.Dish
	Err string
}

func GetDishes (w http.ResponseWriter, r *http.Request) {
	conn := db.Connect()
	dishes, err := models.GetDishes(conn)
	if err != nil {
		err = getDishesTmpl.ExecuteTemplate(w, "base", DishesTemplateData{Dishes: nil, Err: err.Error()})
		return
	}
	err = getDishesTmpl.ExecuteTemplate(w, "base", &DishesTemplateData{Dishes: dishes, Err: ""})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
