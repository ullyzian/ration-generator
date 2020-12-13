package routes

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("./pkg/static"))
	r.Handle("/static/", http.StripPrefix("/static/", fs))
	// routes
	r.HandleFunc("/", handleRoot(s))
	//s.router.HandleFunc("/dishes", s.handleGetDishes())
	//s.router.HandleFunc("/dishes/create", s.handleCreateDish())
	//s.router.HandleFunc("/dishes/{id}", s.handleEditDish())
	//s.router.HandleFunc("/programs", s.handleGetPrograms())
	//s.router.HandleFunc("/programs/create", s.handleCreateProgram())
	return r
}


