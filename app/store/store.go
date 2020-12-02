package store

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Store struct {
	config *Config
	db     *sql.DB
	dishesRepository *DishesRepository
	programRepository *ProgramsRepository
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}
	s.db = db
	return nil
}

func (s *Store) Close() error {
	if err:=s.db.Close(); err != nil {
		return err
	}
	return nil
}

func (s *Store) Dish() *DishesRepository {
	if s.dishesRepository != nil {
		return s.dishesRepository
	}

	s.dishesRepository = &DishesRepository{
		store: s,
	}

	return s.dishesRepository
}

func (s *Store) Program() *ProgramsRepository {
	if s.programRepository != nil {
		return s.programRepository
	}

	s.programRepository = &ProgramsRepository{
		store: s,
	}

	return s.programRepository
}
