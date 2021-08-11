package apiserver

import (
	database "backend/Database"
	"backend/dbModels"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type APIServer struct {
	config   *Config
	logger   *logrus.Logger
	router   *mux.Router
	database *database.DataBase
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.logger.Info("Start server")

	s.configureRouter()

	if err := s.configureDataBase(); err != nil {
		return err
	}

	acc := s.database.Account()
	test := &dbModels.Account{
		Nick:         "TestNick123",
		Hashpassword: "123123",
	}
	acc.Create(test)
	return http.ListenAndServe(s.config.BingAddr, s.router)
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *APIServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}

func (s *APIServer) configureDataBase() error {
	st := database.New(s.config.database)
	if err := st.Open(); err != nil {
		return err
	}
	s.database = st

	return nil
}
