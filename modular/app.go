package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	requestid "github.com/aphistic/negroni-requestid"
	"github.com/calazans10/go-structure-examples/modular/movies"
	"github.com/calazans10/go-structure-examples/modular/reviews"
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
	moviesRepository := movies.NewPostgresRepository(a.DB)
	moviesService := movies.NewService(moviesRepository)
	reviewsRepository := reviews.NewPostgresRepository(a.DB)
	reviewsService := reviews.NewService(reviewsRepository)

	a.Router = mux.NewRouter()
	a.Router.Handle("/movies", movies.GetMovies(moviesService)).Methods("GET")
	a.Router.Handle("/movies/{id}", CheckMovieID(movies.GetMovie(moviesService))).Methods("GET")
	a.Router.Handle("/movies/{id}/reviews", CheckMovieID(reviews.GetMovieReviews(reviewsService))).Methods("GET")
	a.Router.Handle("/movies", movies.AddMovie(moviesService)).Methods("POST")
	a.Router.Handle("/movies/{id}/reviews", CheckMovieID(reviews.AddMovieReview(reviewsService))).Methods("POST")
}
