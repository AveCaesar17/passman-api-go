package store_test

import (
	"os"
	"testing"
)

var (
	DatabaseURL string
)

func TestMain(m *testing.M) {
	DatabaseURL = "host=45.142.212.246 dbname=apiserver sslmode=disable user=admin password=caesar"
	os.Exit(m.Run())

}
