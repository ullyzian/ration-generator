package server

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/ullyzian/ration-generator/app/store"
	"net/http"
)

type Server struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store *store.Store
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
	fs := http.FileServer(http.Dir("./app/static"))
	s.router.Handle("/static/", http.StripPrefix("/static/", fs))
	// routes
	s.router.HandleFunc("/", s.handleRoot())
	s.router.HandleFunc("/dishes", s.handleGetDishes())
	s.router.HandleFunc("/dishes/create", s.handleCreateDish())
	s.router.HandleFunc("/programs", s.handleGetPrograms())
	s.router.HandleFunc("/programs/create", s.handleCreateProgram())
}

func (s *Server) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st
	return nil
}

