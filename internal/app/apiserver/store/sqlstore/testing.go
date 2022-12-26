package sqlstore

import (
	"database/sql"
	"strings"
	"testing"
)

func TestDB(test *testing.T, databaseURL string, databaseDbName string, databaseUser string, databasePassword string) (*sql.DB, func(...string)) {
	test.Helper()
	url := "host=" + databaseURL + " dbname=" + databaseDbName + " sslmode=disable user=" + databaseUser + " password=" + databasePassword
	db, err := sql.Open("postgres", url)
	if err != nil {
		test.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		test.Fatal(err)
	}
	return db, func(tables ...string) {
		if len(tables) > 0 {
			db.Exec("TRUNCATE %S CASCADE", strings.Join(tables, ", "))
		}
		db.Close()
	}
}
