package store

import (
	"fmt"
	"strings"
	"testing"
)

func TestStore(test *testing.T, databaseURL string, databaseDbName string, databaseUser string, databasePassword string) (*Store, func(...string)) {
	test.Helper()
	config := NewConfig()
	config.DatabaseURL = databaseURL
	config.DatabaseDBname = databaseDbName
	config.DatabaseUser = databaseUser
	config.DatabasePassword = databasePassword
	s := New(config)
	if err := s.Open(); err != nil {
		fmt.Println("debug")
		fmt.Println("debug")
		test.Fatal(err)
	}
	return s, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := s.db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", "))); err != nil {
				test.Fatal(err)
			}
		}
	}
}
