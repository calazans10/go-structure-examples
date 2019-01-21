package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	requestid "github.com/aphistic/negroni-requestid"
	"github.com/calazans10/go-structure-examples/layered/handlers"
	"github.com/calazans10/go-structure-examples/layered/models"
	"github.com/calazans10/go-structure-examples/layered/storage"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	negronilogrus "github.com/meatballhat/negroni-logrus"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

// App defines the app structure
type App struct {
	DB     *sqlx.DB
	Router *mux.Router
}

// Initialize creates the DB connection and prepares all the routes
func (a *App) Initialize(user, password, dbname string) {
	a.initializeDatabase(user, password, dbname)
	a.initializeRoutes()
}

// Run initializes the server
func (a *App) Run(appEnv, appPort string) {
	c := cors.AllowAll()

	n := negroni.New()
	n.Use(negroni.NewRecovery())
	n.Use(negronilogrus.NewCustomMiddleware(logrus.InfoLevel, &logrus.JSONFormatter{TimestampFormat: time.RFC3339Nano}, "web"))
	n.Use(c)
	n.Use(requestid.NewMiddleware())
	n.UseHandler(a.Router)

	log.Printf("Starting server in %s mode", appEnv)
	log.Fatal(http.ListenAndServe(":8000", n))
}

func (a *App) initializeDatabase(user, password, dbname string) {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)
	db, err := sqlx.Connect("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}
	a.DB = db

	seed := NewSeed(a.DB)
	seed.Pollute()
}

func (a *App) initializeRoutes() {
	repo := storage.NewPostgresRepository(a.DB)
	service := models.NewService(repo)

	a.Router = mux.NewRouter()
	a.Router.Handle("/movies", handlers.GetMovies(service)).Methods("GET")
	a.Router.Handle("/movies/{id}", CheckMovieID(handlers.GetMovie(service))).Methods("GET")
	a.Router.Handle("/movies/{id}/reviews", CheckMovieID(handlers.GetMovieReviews(service))).Methods("GET")
	a.Router.Handle("/movies", handlers.AddMovie(service)).Methods("POST")
	a.Router.Handle("/movies/{id}/reviews", CheckMovieID(handlers.AddMovieReview(service))).Methods("POST")
}
