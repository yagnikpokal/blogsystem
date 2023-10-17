package api

import (
	"backend/api/mocks" // Import the generated mock package
	"backend/pkg/models"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
)

// Table driven test and gomock package to generate the mock data
func TestAllArticle(t *testing.T) {
	testCases := []struct {
		name                    string
		mockDBAllArticlesReturn []*models.Articles // The value to be returned by the mockDB.AllArticles() call
		expectedStatusCode      int
	}{
		{
			name: "Success Case",
			mockDBAllArticlesReturn: []*models.Articles{
				&models.Articles{ID: 1, Title: "Article 1", Content: "Content 1", Author: "Author 1"},
				&models.Articles{ID: 2, Title: "Article 2", Content: "Content 2", Author: "Author 2"},
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:                    "No Articles Found",
			mockDBAllArticlesReturn: []*models.Articles{}, // No articles returned
			expectedStatusCode:      http.StatusNotFound,
		},
	}

	// Iterate over the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new instance of the gomock controller
			ctrl := gomock.NewController(t)
			defer ctrl.Finish() // Ensure the controller is closed when the test ends

			// Create mock objects for your dependencies
			mockDB := mocks.NewMockDBInterface(ctrl)
			mockUtility := mocks.NewMockUtilityInterface(ctrl)

			// Create an instance of your Application with the mock dependencies
			app := Application{
				DB:      mockDB,
				Utility: mockUtility,
			}

			// Set expectations for the mockDB's AllArticles method
			mockDB.EXPECT().AllArticles().Return(tc.mockDBAllArticlesReturn, nil)

			// Set expectations for the mockUtility's WriteJSON method
			//			mockUtility.EXPECT().WriteJSON(gomock.Any(), tc.expectedStatusCode, gomock.Any()).Return(nil)

			// Create an HTTP response recorder and request for testing
			w := httptest.NewRecorder()
			r, _ := http.NewRequest("GET", "/articles", nil)

			// Call the function to be tested
			app.AllArticle(w, r)

			// Add your assertions here to check the response status code
			if w.Code != tc.expectedStatusCode {
				t.Errorf("Expected status code %d, but got %d", tc.expectedStatusCode, w.Code)
			}
		})
	}
}
func TestInsertArticle_ErrorCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks.NewMockDBInterface(ctrl)
	mockUtility := mocks.NewMockUtilityInterface(ctrl)

	app := Application{
		DB:      mockDB,
		Utility: mockUtility,
	}

	// Create a sample article for your test
	sampleArticle := &models.Articles{
		ID:      1,
		Title:   "Sample Article",
		Content: "Sample Content",
		Author:  "Sample Author",
	}

	// Simulate a database error
	expectedDBError := errors.New("Some database error")

	// Set expectations for the mockDB's CreateArticle method to return an error
	mockDB.EXPECT().CreateArticle(gomock.Any()).Return(0, expectedDBError)

	// Set expectations for the mockUtility's WriteJSON method to return a 404 status code
	//	mockUtility.EXPECT().WriteJSON(gomock.Any(), http.StatusNotFound, gomock.Any()).Return(nil)

	w := httptest.NewRecorder()

	// Create a request with the sample article as the JSON body
	body, err := json.Marshal(sampleArticle)
	if err != nil {
		t.Fatal("Failed to marshal JSON:", err)
	}

	r, _ := http.NewRequest("POST", "/articles", bytes.NewBuffer(body))
	r.Header.Set("Content-Type", "application/json")

	app.InsertArticle(w, r)
}
func TestGetArticle_DatabaseError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := mocks.NewMockDBInterface(ctrl)
	mockUtility := mocks.NewMockUtilityInterface(ctrl)

	app := Application{
		DB:      mockDB,
		Utility: mockUtility,
	}

	// Set expectations for the mockDB's OneArticle method to return a database error
	expectedDBError := errors.New("Database error")
	//articleID := 1 // Use a valid article ID

	mockDB.EXPECT().OneArticle(0).Return(nil, expectedDBError)

	// Set expectations for the mockUtility's WriteJSON method to return a 500 status code
	// mockUtility.EXPECT().WriteJSON(gomock.Any(), http.StatusInternalServerError, gomock.Any()).Return(nil)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/articles/1", nil)
	chiURLParam := chi.NewRouteContext()
	chiURLParam.URLParams.Add("id", "1") // Correctly set the article ID
	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, chiURLParam))

	app.GetArticle(w, r)
}
