package teststore_test

import (
	"testing"

	"github.com/AveCaesar17/basic-server-go.git/internal/app/apiserver/store"
	"github.com/AveCaesar17/basic-server-go.git/internal/app/apiserver/store/teststore"
	"github.com/AveCaesar17/basic-server-go.git/internal/app/model"
	"github.com/stretchr/testify/assert"
)

func return_config() (DatabaseURL string, DatabaseDbname string, DatabaseUser string, DatabasePassword string) {
	DatabaseURL = "45.142.212.246"
	DatabaseDBName := "apiserver"
	DatabaseUser = "admin"
	DatabasePassword = "caesar"

	return DatabaseURL, DatabaseDBName, DatabaseUser, DatabasePassword
}

func TestUserRepository(t *testing.T) {

	s := teststore.New()
	user := model.TestUser((t))

	assert.NoError(t, s.User().Create(model.TestUser(t)))
	assert.NotNil(t, user)
}
func TestUserRepository_FindByUsername(t *testing.T) {

	s := teststore.New()
	username := "caesar"
	_, err := s.User().FindByUsername(username)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.User().Create(&model.User{
		Username: "caesar",
		Pub_Key:  "pubkey_test",
	})
	u, err := s.User().FindByUsername(username)
	assert.NoError(t, err)
	assert.NotNil(t, u)

}
