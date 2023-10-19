package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"backend/mocks" // Import the generated mock package

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/golang/mock/gomock" // Import gomock
	"github.com/stretchr/testify/assert"
)

func TestRoutes(t *testing.T) {
	// Create a new instance of the mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock instance of the Application
	mockApp := mocks.NewMockRoutes(ctrl)

	// Create a request to test the routes
	req := httptest.NewRequest("GET", "/", nil)

	// Create a recorder to capture the response
	recorder := httptest.NewRecorder()

	// Initialize the router and set up routes
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Get("/articles", mockApp.AllArticle)
	router.Get("/articles/{id}", mockApp.GetArticle)
	router.Post("/articles", mockApp.InsertArticle)

	// Serve the request
	router.ServeHTTP(recorder, req)
}

func TestRoutes_Request(t *testing.T) {
	// Create a new instance of the Application
	app := &Application{}

	// Create a request to test the routes
	req := httptest.NewRequest("GET", "/", nil)

	// Create a recorder to capture the response
	recorder := httptest.NewRecorder()

	// Initialize the router and set up routes
	router := app.Routes()

	// Serve the request
	router.ServeHTTP(recorder, req)

	// Check the response status code and other assertions
	assert.Equal(t, http.StatusOK, recorder.Code)
}

// Unit test using table driven test
func TestRoutes_Healthcheck(t *testing.T) {
	testCases := []struct {
		name         string
		method       string
		path         string
		expectedCode int
	}{
		{
			name:         "Successful GET request to Healthcheck",
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
			assert.Equal(t, tc.expectedCode, recorder.Code)
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

			assert.Equal(t, tc.expectedCode, recorder.Code)
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
			name:         "Negative test case for InsertArticle",
			method:       "POST",
			path:         "/articles",
			expectedCode: 400,
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
