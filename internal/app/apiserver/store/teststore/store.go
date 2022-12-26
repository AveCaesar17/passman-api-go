package teststore

import (
	"github.com/AveCaesar17/basic-server-go.git/internal/app/apiserver/store"
	"github.com/AveCaesar17/basic-server-go.git/internal/app/model"
	_ "github.com/lib/pq"
)

type Store struct {
	UserRepository *UserRepository
}

func New() *Store {
	return &Store{}

}

func (s *Store) User() store.UserRepository {
	if s.UserRepository != nil {
		return s.UserRepository
	}
	s.UserRepository = &UserRepository{
		store: s,
		users: make(map[string]*model.User),
	}
	return s.UserRepository
}

func GetMeCred(DatabaseURL string, DatabaseDBname string, DatabaseUser string, DatabasePassword string) (string, error) {
	url := "host=" + DatabaseURL + " dbname=" + DatabaseDBname + " sslmode=disable user=" + DatabaseUser + " password=" + DatabasePassword
	return url, nil

}
