package apiserver

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/AveCaesar17/basic-server-go.git/internal/app/apiserver/store"
	"github.com/AveCaesar17/basic-server-go.git/internal/app/model"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var (
	errIncorrectUsernameOrPassword = errors.New("incorrect username or password")
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
	s.router.HandleFunc("/create_user", s.handleUserCreate()).Methods("POST")
	s.router.HandleFunc("/session", s.handleSessionCreate()).Methods("POST")
}
func (s *server) handleUserCreate() http.HandlerFunc {
	type request struct {
		Username string `json:"username"`
		PubKey   string `json:"pubkey"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		u := &model.User{
			Username: req.Username,
			Pub_Key:  req.PubKey,
		}

		if err := s.store.User().Create(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
		}

		u.Sanitize()
		s.response(w, r, http.StatusCreated, u)
	}
}
func (s *server) handleSessionCreate() http.HandlerFunc {
	type request struct {
		Username string `json:"username"`
		PubKey   string `json:"Pubkey"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		u, err := s.store.User().FindByUsername(req.Username)
		if err != nil || !u.CheckPassword(req.PubKey) {
			s.error(w, r, http.StatusUnauthorized, errIncorrectUsernameOrPassword)
			return
		}
		s.response(w, r, http.StatusOK, nil)
	}
}
func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.response(w, r, code, map[string]string{"error": err.Error()})
}
func (s *server) response(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
