package store

import "github.com/AveCaesar17/basic-server-go.git/internal/app/model"

type UserRepository interface {
	Create(*model.User) error
	FindByUsername(string) (*model.User, error)
}
