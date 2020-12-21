package server

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/ullyzian/ration-generator/pkg/store"
	"net/http"
)

type Server struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

func New(config *Config) *Server {
	return &Server{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *Server) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return err
	}

	s.logger.Info("Starting web server")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *Server) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *Server) configureRouter() {
	// file server
	fs := http.FileServer(http.Dir("./pkg/static"))
	s.router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	// routes
	s.router.HandleFunc("/", s.handleRoot())
	s.router.HandleFunc("/generator", s.handleGenerator())
	s.router.HandleFunc("/nutrition", s.handleProgramsList())
	s.router.HandleFunc("/dishes", s.handleGetDishes())
	s.router.HandleFunc("/dishes/create", s.handleCreateDish())
	s.router.HandleFunc("/dishes/{id}/edit", s.handleEditDish())
	s.router.HandleFunc("/dishes/{id}/delete", s.handleDeleteDish())
	s.router.HandleFunc("/programs", s.handleGetPrograms())
	s.router.HandleFunc("/programs/create", s.handleCreateProgram())
	s.router.HandleFunc("/programs/{id}/edit", s.handleEditProgram())
	s.router.HandleFunc("/programs/{id}/delete", s.handleDeleteProgram())
}

func (s *Server) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st
	return nil
}
