package apiserver

import (
	"net/http"

	"github.com/AveCaesar17/basic-server-go.git/internal/app/apiserver/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type server struct {
	router *mux.Router
	loger  *logrus.Logger
	store  store.Store
}

func newServer(store store.Store) *server {
	serv := &server{
		router: mux.NewRouter(),
		loger:  logrus.New(),
		store:  store,
	}
	serv.configureRouter()

	return serv
}
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
func (s *server) configureRouter() {
	s.router.HandleFunc("/test", s.handleUserCreate()).Methods("POST")
}
func (s *server) handleUserCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
