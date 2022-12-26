package model_test

import (
	"testing"

	"github.com/AveCaesar17/basic-server-go.git/internal/app/model"
	"github.com/stretchr/testify/assert"
)

func TestUser_BeforeCreate(test *testing.T) {
	testCases := []struct {
		name    string
		user    func() *model.User
		isValid bool
	}{
		{
			name: "valid",
			user: func() *model.User {
				return model.TestUser(test)
			},
			isValid: true,
		},
		{
			name: "empotyUsername",
			user: func() *model.User {
				user := model.TestUser(test)
				user.Username = ""
				return user
			},
			isValid: false,
		},
		{
			name: "empotyPassword",
			user: func() *model.User {
				user := model.TestUser(test)
				user.Pub_Key = ""
				return user
			},
			isValid: false,
		},
	}
	for _, tc := range testCases {
		test.Run(tc.name, func(test *testing.T) {
			if tc.isValid {
				assert.NoError(test, tc.user().Validate())
			} else {
				assert.Error(&testing.F{}, tc.user().Validate())
			}
		})

	}
}

func TestUserValidate(test *testing.T) {
	user := model.TestUser(test)
	assert.NoError(test, user.Validate())
}
