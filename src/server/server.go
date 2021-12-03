package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/juricaKenda/golang-rest-example/src/model"
	"net/http"
)

type Server struct {
	host   string
	port   int64
	router *mux.Router

	articleRepo IArticleRepository
	log         ILogger
	notif       Notif
}

type IArticleRepository interface {
	GetAllArticles() ([]model.Article, error)
}

type ILogger interface {
	Error(msg string)
	Trace(msg string)
}

type Notif interface {
	Notify(msg string)
}

func New(host string, port int64, articles IArticleRepository, log ILogger, notif Notif) *Server {
	s := &Server{
		host:        host,
		port:        port,
		router:      mux.NewRouter(),
		articleRepo: articles,
		log:         log,
		notif:       notif,
	}
	s.registerRoutes()
	return s
}

func (s *Server) Start() error {
	return http.ListenAndServe(s.accessPoint(), s.router)
}

func (s *Server) accessPoint() string {
	return fmt.Sprintf("%s:%d", s.host, s.port)
}

func (s *Server) registerRoutes() {
	s.log.Trace("registering router...")
	s.router.HandleFunc("/articles", s.getAllArticles()).Methods("GET")
}
