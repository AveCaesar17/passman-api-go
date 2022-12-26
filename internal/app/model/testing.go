package model

import "testing"

func TestUser(test *testing.T) *User {
	return &User{
		Username: "caesar",
		Pub_Key:  "test_pubkey",
	}
}
