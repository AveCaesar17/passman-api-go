package apiserver

import (
	"fmt"
	"io"
	"net/http"

	"github.com/AveCaesar17/basic-server-go.git/internal/app/apiserver/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (server *APIServer) Start() error {
	if err := server.configureLogger(); err != nil {
		return err
	}
	server.configureRouter()
	if err := server.configureStore(); err != nil {
		return err
	}
	fmt.Println("Server is started...")
	return http.ListenAndServe(server.config.BindAddr, server.router)
}

func (server *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(server.config.LogLevel)
	if err != nil {
		return err
	}
	server.logger.SetLevel(level)

	return nil
}
func (server *APIServer) configureRouter() {
	server.router.HandleFunc("/hello", server.handleHello())
	server.router.HandleFunc("/test", server.handleTest())
}
func (server *APIServer) configureStore() error {
	st := store.New(server.config.Store)
	if err := st.Open(); err != nil {
		return err
	}
	server.store = st
	return nil

}

func (server *APIServer) handleHello() http.HandlerFunc {

	type request struct {
		name string
	}
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hi")
	}
}
func (serevr *APIServer) handleTest() http.HandlerFunc {

	type request struct {
		name string
	}
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "test")
	}
}
