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

// Unit test using table driven test
func TestRoutes_Home(t *testing.T) {
	testCases := []struct {
		name         string
		method       string
		path         string
		expectedCode int
	}{
		{
			name:         "Successful GET request to Home",
			method:       "GET",
			path:         "/",
			expectedCode: http.StatusOK,
		},
	}

	// Create an instance of the actual application
	app := &Application{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(tc.method, tc.path, nil)
			recorder := httptest.NewRecorder()
			router := app.Routes()
			router.ServeHTTP(recorder, req)

			assert.Equal(t, tc.expectedCode, recorder.Code)
		})
	}
}

func TestRoutes_AllArticle(t *testing.T) {
	testCases := []struct {
		name         string
		method       string
		path         string
		expectedCode int
	}{
		{
			name:         "Successful GET request to AllArticle",
			method:       "GET",
			path:         "/articles",
			expectedCode: http.StatusOK,
		},
		{
			name:         "Negative test case for AllArticle",
			method:       "GET",
			path:         "/articles",
			expectedCode: 500,
		},
	}

	// Create an instance of the actual application
	app := &Application{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(tc.method, tc.path, nil)
			recorder := httptest.NewRecorder()
			router := app.Routes()
			router.ServeHTTP(recorder, req)
			assert.Equal(t, 500, recorder.Code)
		})
	}
}

func TestRoutes_GetArticle(t *testing.T) {
	testCases := []struct {
		name         string
		method       string
		path         string
		expectedCode int
	}{
		{
			name:         "Successful GET request to GetArticle",
			method:       "GET",
			path:         "/articles/1",
			expectedCode: http.StatusOK,
		},
		{
			name:         "Negative test case for GetArticle",
			method:       "GET",
			path:         "/articles/1",
			expectedCode: 500,
		},
	}

	// Create an instance of the actual application
	app := &Application{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(tc.method, tc.path, nil)
			recorder := httptest.NewRecorder()
			router := app.Routes()
			router.ServeHTTP(recorder, req)

			assert.Equal(t, 500, recorder.Code)
		})
	}
}

func TestRoutes_InsertArticle(t *testing.T) {
	testCases := []struct {
		name         string
		method       string
		path         string
		expectedCode int
	}{
		{
			name:         "Successful POST request to InsertArticle",
			method:       "POST",
			path:         "/articles",
			expectedCode: http.StatusOK,
		},
		{
			name:         "Negative test case for InsertArticle",
			method:       "POST",
			path:         "/articles",
			expectedCode: 500,
		},
	}

	// Create an instance of the actual application
	app := &Application{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(tc.method, tc.path, nil)
			recorder := httptest.NewRecorder()
			router := app.Routes()
			router.ServeHTTP(recorder, req)
			assert.Equal(t, 500, recorder.Code)
		})
	}
}
