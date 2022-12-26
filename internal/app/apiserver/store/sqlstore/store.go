package sqlstore

import (
	"database/sql"

	"github.com/AveCaesar17/basic-server-go.git/internal/app/apiserver/store"
	_ "github.com/lib/pq"
)

type Store struct {
	db             *sql.DB
	UserRepository *UserRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) User() store.UserRepository {
	if s.UserRepository != nil {
		return s.UserRepository
	}
	s.UserRepository = &UserRepository{
		store: s,
	}
	return s.UserRepository
}

func GetMeCred(DatabaseURL string, DatabaseDBname string, DatabaseUser string, DatabasePassword string) (string, error) {
	url := "host=" + DatabaseURL + " dbname=" + DatabaseDBname + " sslmode=disable user=" + DatabaseUser + " password=" + DatabasePassword
	return url, nil

}
