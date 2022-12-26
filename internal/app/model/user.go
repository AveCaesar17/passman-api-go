package model

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID               int
	Username         string
	Encrypt_Password string
	Pub_Key          string
}

func (user *User) Validate() error {
	return validation.ValidateStruct(
		user,
		validation.Field(&user.Username, validation.Required, is.ASCII),
		validation.Field(&user.Pub_Key, validation.Required, is.ASCII),
	)
}
func (user *User) BeforeCreate() error {
	if len(user.Pub_Key) > 0 {
		enc, err := EncryptString(user.Pub_Key)
		if err != nil {
			return err
		}

		user.Encrypt_Password = enc

	}
	return nil

}
func EncryptString(str string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
