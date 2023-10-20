package api

import (
	"backend/pkg/repository/dbrepo"
	services "backend/services/articles"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	_ "github.com/swaggo/http-swagger/example/go-chi/docs"
)

type Application struct {
	DSN            string
	DB             dbrepo.DatabaseRepo
	Utility        UtilityInterface
	ArticleService *services.ArticleService
}
type Routes interface {
	HealthCheck(w http.ResponseWriter, r *http.Request)
	AllArticle(w http.ResponseWriter, r *http.Request)
	GetArticle(w http.ResponseWriter, r *http.Request)
	InsertArticle(w http.ResponseWriter, r *http.Request)
}

func (app *Application) Routes() http.Handler {
	// create a router mux
	mux := chi.NewRouter()
	// Create a new CORS middleware instance with your desired options.
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Replace with allowed origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           600,
	})
	// Use the CORS middleware
	mux.Use(c.Handler)
	mux.Use(middleware.Recoverer)
	mux.Get("/", app.HealthCheck)
	mux.Get("/articles", app.AllArticle)
	mux.Get("/articles/{id}", app.GetArticle)
	mux.Post("/articles", app.InsertArticle)

	return mux
}
