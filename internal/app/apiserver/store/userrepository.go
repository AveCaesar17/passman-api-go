package store

import (
	"github.com/AveCaesar17/basic-server-go.git/internal/app/model"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(user *model.User) (*model.User, error) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.BeforeCreate(); err != nil {
		return nil, err
	}
	if err := r.store.db.QueryRow("INSERT INTO users (username,pub_key) VALUES ($1,$2) RETURNING id", user.Username, user.Pub_Key).Scan(&user.ID); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) FindByUsername(username string) (*model.User, error) {
	user := &model.User{}
	if err := r.store.db.QueryRow("SELECT id, username, pub_key FROM users WHERE username = $1", username).Scan(&user.ID, &user.Username, &user.Pub_Key); err != nil {
		return nil, err
	}
	return user, nil
}
