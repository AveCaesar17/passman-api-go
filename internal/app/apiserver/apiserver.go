package apiserver

import (
	"database/sql"
	"net/http"

	"github.com/AveCaesar17/basic-server-go.git/internal/app/apiserver/store/sqlstore"
)

func Start(config *Config) error {
	url := "host=" + config.DatabaseURL + " dbname=" + config.DatabaseDBName + " sslmode=disable user=" + config.DatabaseUser + " password=" + config.DatabasePassword

	db, err := newDB(url)
	if err != nil {
		return err
	}
	defer db.Close()
	store := sqlstore.New(db)
	serv := newServer(store)
	return http.ListenAndServe(config.BindAddr, serv)
}
func newDB(url string) (*sql.DB, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
