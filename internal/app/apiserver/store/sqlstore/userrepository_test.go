package sqlstore_test

import (
	"fmt"
	"testing"

	"github.com/AveCaesar17/basic-server-go.git/internal/app/apiserver/store"
	"github.com/AveCaesar17/basic-server-go.git/internal/app/apiserver/store/sqlstore"
	"github.com/AveCaesar17/basic-server-go.git/internal/app/model"
	"github.com/stretchr/testify/assert"
)

func return_config() (DatabaseURL string, DatabaseDbname string, DatabaseUser string, DatabasePassword string) {
	DatabaseURL = ""
	DatabaseDBName := ""
	DatabaseUser = ""
	DatabasePassword = ""

	return DatabaseURL, DatabaseDBName, DatabaseUser, DatabasePassword
}

func TestUserRepository(t *testing.T) {
	DatabaseURL, DatabaseDbName, DatabaseUser, DatabasePassword := return_config()
	db, teardown := sqlstore.TestDB(t, DatabaseURL, DatabaseDbName, DatabaseUser, DatabasePassword)
	defer teardown("users")
	s := sqlstore.New(db)

	user := model.TestUser(t)
	fmt.Println(user)

	assert.NoError(t, s.User().Create(user))
	assert.NotNil(t, user)
}
func TestUserRepository_FindByUsername(t *testing.T) {
	DatabaseURL, DatabaseDbName, DatabaseUser, DatabasePassword := return_config()

	db, teardown := sqlstore.TestDB(t, DatabaseURL, DatabaseDbName, DatabaseUser, DatabasePassword)
	defer teardown("users")
	s := sqlstore.New(db)
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
