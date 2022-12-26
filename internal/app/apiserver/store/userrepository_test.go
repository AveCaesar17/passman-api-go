package store_test

import (
	"testing"

	"github.com/AveCaesar17/basic-server-go.git/internal/app/apiserver/store"
	"github.com/AveCaesar17/basic-server-go.git/internal/app/model"
	"github.com/stretchr/testify/assert"
)

func return_config() (DatabaseURL string, DatabaseDbname string, DatabaseUser string, DatabasePassword string) {
	DatabaseURL = ""
	DatabaseDBName := "apiserver"
	DatabaseUser = "admin"
	DatabasePassword = ""

	return DatabaseURL, DatabaseDBName, DatabaseUser, DatabasePassword
}

func TestUserRepository(t *testing.T) {
	DatabaseURL, DatabaseDbName, DatabaseUser, DatabasePassword := return_config()
	s, teardown := store.TestStore(t, DatabaseURL, DatabaseDbName, DatabaseUser, DatabasePassword)
	defer teardown("users")
	u, err := s.User().Create(&model.User{
		Username: "caesar",
		Pub_Key:  "pubkey_test",
	})
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
func TestUserRepository_FindByUsername(t *testing.T) {
	DatabaseURL, DatabaseDbName, DatabaseUser, DatabasePassword := return_config()

	s, teardown := store.TestStore(t, DatabaseURL, DatabaseDbName, DatabaseUser, DatabasePassword)
	defer teardown("users")
	username := "caesar"
	_, err := s.User().FindByUsername(username)
	assert.Error(t, err)
	s.User().Create(&model.User{
		Username: "caesar",
		Pub_Key:  "pubkey_test",
	})
	u, err := s.User().FindByUsername(username)
	assert.NoError(t, err)
	assert.NotNil(t, u)

}
