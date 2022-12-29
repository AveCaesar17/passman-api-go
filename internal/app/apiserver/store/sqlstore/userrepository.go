package sqlstore

import (
	"database/sql"
	"fmt"

	"github.com/AveCaesar17/basic-server-go.git/internal/app/apiserver/store"
	"github.com/AveCaesar17/basic-server-go.git/internal/app/model"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(user *model.User) error {
	if err := user.Validate(); err != nil {
		return err
	}
	fmt.Println(user.Username, user.Pub_Key)
	return r.store.db.QueryRow(

		"INSERT INTO users (username, pub_key) VALUES ($1,$2) RETURNING id",
		user.Username, user.Pub_Key,
	).Scan(&user.ID)
}

func (r *UserRepository) FindByUsername(username string) (*model.User, error) {
	user := &model.User{}
	if err := r.store.db.QueryRow("SELECT id, username, pub_key FROM users WHERE username = $1", username).Scan(&user.ID, &user.Username, &user.Pub_Key); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return user, nil
}
