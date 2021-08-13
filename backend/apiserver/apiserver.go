package apiserver

import (
	database "backend/Database"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type APIServer struct {
	config   *Config
	logger   *logrus.Logger
	router   *mux.Router
	database *database.DataBase
}

//configuration server
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

//func to configuration modules and start server
func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.logger.Info("Start server")
	s.configureRouter()
	if err := s.configureDataBase(); err != nil {
		return err
	}
	return http.ListenAndServe(s.config.BingAddr, s.router)
}

//configuration logger
func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

// configuration api
func (s *APIServer) configureRouter() {

	s.router.HandleFunc("/api/user/login", s.Login()).Methods("POST")
	//s.router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	//s.router.HandleFunc("/api/contacts/new", controllers.CreateContact).Methods("POST")
	//s.router.HandleFunc("/api/me/contacts", controllers.GetContactsFor).Methods("GET") //  user/2/contacts

	//s.router.Use(app.JwtAuthentication) //attach JWT auth middleware

	s.logger.Info("4000")

	// err := http.ListenAndServe(":4000", s.router) //Launch the app, visit localhost:8000/api
	// if err != nil {
	// 	s.logger.Info(err)
	// }
}

//configuration data base
func (s *APIServer) configureDataBase() error {
	st := database.New(s.config.database)
	if err := st.Open(); err != nil {
		return err
	}
	s.database = st

	return nil
}
