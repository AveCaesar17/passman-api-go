package teststore

import (
	"github.com/AveCaesar17/basic-server-go.git/internal/app/apiserver/store"
	"github.com/AveCaesar17/basic-server-go.git/internal/app/model"
)

type UserRepository struct {
	store *Store
	users map[string]*model.User
}

func (r *UserRepository) Create(user *model.User) error {
	if err := user.Validate(); err != nil {
		return err
	}
	r.users[user.Username] = user
	user.ID = len(r.users)
	return nil
}
func (r *UserRepository) FindByUsername(Username string) (*model.User, error) {
	user, ok := r.users[Username]
	if !ok {
		return nil, store.ErrRecordNotFound
	}
	return user, nil
}
