package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/stretchr/testify/assert"
)

func TestRoutes(t *testing.T) {
	// Create a new instance of the mock application
	mockApp := &mockApplication{}

	// Create a request to test the routes
	req := httptest.NewRequest("GET", "/", nil)

	// Create a recorder to capture the response
	recorder := httptest.NewRecorder()

	// Initialize the router and set up routes
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Get("/", mockApp.Home)
	router.Get("/articles", mockApp.AllArticle)
	router.Get("/articles/{id}", mockApp.GetArticle)
	router.Post("/articles", mockApp.InsertArticle)

	// Serve the request
	router.ServeHTTP(recorder, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, recorder.Code)

	// You can add more test cases for other routes as needed
}

// Mock application for testing
type mockApplication struct {
	// Define mock methods here\

}

func (ma *mockApplication) Home(w http.ResponseWriter, r *http.Request) {
	// Implement mock behavior for Home
}

func (ma *mockApplication) AllArticle(w http.ResponseWriter, r *http.Request) {
	// Implement mock behavior for AllArticle
}

func (ma *mockApplication) GetArticle(w http.ResponseWriter, r *http.Request) {
	// Implement mock behavior for GetArticle
}

func (ma *mockApplication) InsertArticle(w http.ResponseWriter, r *http.Request) {
	// Implement mock behavior for InsertArticle
}
func TestRoutes_Home(t *testing.T) {
	// Create an instance of the actual application
	app := &Application{}

	// Create a request for the home route
	req := httptest.NewRequest("GET", "/", nil)

	// Create a recorder to capture the response
	recorder := httptest.NewRecorder()

	// Initialize the router and set up routes
	router := app.Routes()

	// Serve the request
	router.ServeHTTP(recorder, req)

	// Check the response status code for the home route
	assert.Equal(t, http.StatusOK, recorder.Code)
}
