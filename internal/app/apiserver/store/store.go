package store

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Store struct {
	config         *Config
	db             *sql.DB
	UserRepository *UserRepository
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	url := "host=" + s.config.DatabaseURL + " dbname=" + s.config.DatabaseDBname + " sslmode=disable user=" + s.config.DatabaseUser + " password=" + s.config.DatabasePassword
	db, err := sql.Open("postgres", url)
	if err != nil {

		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	s.db = db
	return nil
}

func (s *Store) Clouse() {
	s.db.Close()

}

func (s *Store) User() *UserRepository {
	if s.UserRepository != nil {
		return s.UserRepository
	}
	s.UserRepository = &UserRepository{
		store: s,
	}
	return s.UserRepository
}
