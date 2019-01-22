package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	requestid "github.com/aphistic/negroni-requestid"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	negronilogrus "github.com/meatballhat/negroni-logrus"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

// Server defines the server structure
type Server struct {
	DB     *sqlx.DB
	Router *mux.Router
}

// Initialize creates the DB connection and prepares all the routes
func (s *Server) Initialize(user, password, dbname string) {
	s.initializeDatabase(user, password, dbname)
	s.initializeRoutes()
}

// Run initializes the server
func (s *Server) Run(appEnv, appPort string) {
	c := cors.AllowAll()

	n := negroni.New()
	n.Use(negroni.NewRecovery())
	n.Use(negronilogrus.NewCustomMiddleware(logrus.InfoLevel, &logrus.JSONFormatter{TimestampFormat: time.RFC3339Nano}, "web"))
	n.Use(c)
	n.Use(requestid.NewMiddleware())
	n.UseHandler(s.Router)

	log.Printf("Starting server in %s mode", appEnv)
	log.Fatal(http.ListenAndServe(":8000", n))
}

func (s *Server) initializeDatabase(user, password, dbname string) {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)
	db, err := sqlx.Connect("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}
	s.DB = db

	seed := NewSeed(s.DB)
	seed.Pollute()
}

func (s *Server) initializeRoutes() {
	repo := NewPostgresRepository(s.DB)
	service := NewService(repo)
	handler := NewHandler(service)

	s.Router = mux.NewRouter()
	s.Router.Handle("/movies", handler.GetMovies()).Methods("GET")
	s.Router.Handle("/movies/{id}", CheckMovieID(handler.GetMovie())).Methods("GET")
	s.Router.Handle("/movies/{id}/reviews", CheckMovieID(handler.GetMovieReviews())).Methods("GET")
	s.Router.Handle("/movies", handler.AddMovie()).Methods("POST")
	s.Router.Handle("/movies/{id}/reviews", CheckMovieID(handler.AddMovieReview())).Methods("POST")
}
