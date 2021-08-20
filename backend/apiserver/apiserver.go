package apiserver

import (
	database "backend/Database"
	"backend/app"
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

	//account
	s.router.HandleFunc("/api/user/login", s.Login()).Methods("POST")
	s.router.HandleFunc("/api/user/registration", s.Registration()).Methods("POST")
	s.router.HandleFunc("/api/user/search", s.Search()).Methods("GET")
	//subs
	s.router.HandleFunc("/api/user/mySub", s.GetMySubscribes()).Methods("GET")
	s.router.HandleFunc("/api/user/subUser", s.GetUserSubscribes()).Methods("GET")
	s.router.HandleFunc("/api/user/sub", s.Subscribe()).Methods("POST")
	s.router.HandleFunc("/api/user/sub", s.Unsubscribe()).Methods("DELETE")

	//post
	s.router.HandleFunc("/api/post/create", s.CreatePost()).Methods("POST")
	s.router.HandleFunc("/api/post/delete", s.DeletePost()).Methods("DELETE")
	//like unlike
	s.router.HandleFunc("/api/post/like", s.LikePost()).Methods("POST")
	s.router.HandleFunc("/api/post/like", s.UnlikePost()).Methods("DELETE")

	//technology
	s.router.HandleFunc("/api/technology", s.GetTechNology()).Methods("GET")

	//feed
	s.router.HandleFunc("/api/myFeed", s.GetMyFeed()).Methods("GET")
	s.router.HandleFunc("/api/post/userById", s.GetPostsUserById()).Methods("GET")

	s.router.Use(app.JwtAuthentication)
	s.logger.Info("4000")
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
