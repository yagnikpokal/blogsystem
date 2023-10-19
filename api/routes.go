package api

import (
	"backend/pkg/repository/dbrepo"
	services "backend/services/articles"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

	mux.Use(middleware.Recoverer)

	mux.Get("/", app.HealthCheck)
	mux.Get("/articles", app.AllArticle)
	mux.Get("/articles/{id}", app.GetArticle)
	mux.Post("/articles", app.InsertArticle)

	return mux
}
