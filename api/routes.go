package api

import (
	"backend/pkg/repository"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Application struct {
	DSN     string
	DB      repository.DatabaseRepo
	Utility UtilityInterface
}

func (app *Application) Routes() http.Handler {
	// create a router mux
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)

	mux.Get("/", app.Home)
	mux.Get("/articles", app.AllArticle)
	mux.Get("/articles/{id}", app.GetArticle)
	mux.Post("/articles", app.InsertArticle)

	return mux
}
