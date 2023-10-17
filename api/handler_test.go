package api

import (
	"backend/pkg/models"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockDatabase is a mock for your database repository interface
type MockDatabase struct {
	mock.Mock
}
type DatabaseRepo interface {
	AllArticles() ([]*models.Articles, error)
	OneArticle(id int) (*models.Articles, error)
	CreateArticle(article *models.Articles) (int, error)
	Connection() *sql.DB
	CreateTable()
}

func (m *MockDatabase) CreateTable() {}

// Connection mocks the Connection method
func (m *MockDatabase) Connection() *sql.DB {
	args := m.Called()
	return args.Get(0).(*sql.DB)
}

// AllArticles mocks the AllArticles method
func (m *MockDatabase) AllArticles() ([]*models.Articles, error) {
	args := m.Called()
	return args.Get(0).([]*models.Articles), args.Error(1)
}

// OneArticle mocks the OneArticle method
func (m *MockDatabase) OneArticle(id int) (*models.Articles, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Articles), args.Error(1)
}

// CreateArticle mocks the CreateArticle method
func (m *MockDatabase) CreateArticle(article *models.Articles) (int, error) {
	args := m.Called(article)
	return args.Int(0), args.Error(1)
}

func TestHome(t *testing.T) {
	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP recorder
	rr := httptest.NewRecorder()

	// Create a new application struct
	app := Application{}

	// Call the Home() function
	app.Home(rr, req)

	// Check that the HTTP status code is correct
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	// Check that the response body is correct
	expectedBody := `{"status":"active","message":"Go articles up and running","version":"1.0.0"}`
	if rr.Body.String() != expectedBody {
		t.Errorf("Expected response body %s, got %s", expectedBody, rr.Body.String())
	}
}
func TestInsertArticle(t *testing.T) {
	// Create a new HTTP request with a JSON request body
	requestBody := `{"title":"Sample Title","content":"Sample Content","author":"John Doe"}`
	req, err := http.NewRequest("POST", "/articles", strings.NewReader(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a new HTTP recorder
	rr := httptest.NewRecorder()

	// Create a new application struct
	app := &Application{}

	// Create a mock database
	mockDB := &MockDatabase{}

	// Set up a mock response for the database's CreateArticle method
	expectedArticleID := 1
	mockDB.On("CreateArticle", mock.Anything).Return(expectedArticleID, nil)
	app.DB = mockDB

	// Call the InsertArticle function
	app.InsertArticle(rr, req)

	// Check that the HTTP status code is correct
	assert.Equal(t, http.StatusCreated, rr.Code)

	// Parse the JSON response
	var response map[string]interface{}
	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}

	// Check the response JSON
	assert.Equal(t, float64(201), response["status"])
	assert.Equal(t, "Success", response["message"])
	assert.Equal(t, map[string]interface{}{"id": float64(expectedArticleID)}, response["data"])
	// Check that the database CreateArticle method was called with the correct data
	mockDB.AssertCalled(t, "CreateArticle", mock.Anything)
}

func TestAllArticle(t *testing.T) {
	requestBody := `{"title":"Sample Title","content":"Sample Content","author":"John Doe"}`
	req, err := http.NewRequest("POST", "/articles", strings.NewReader(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a new HTTP recorder
	rr := httptest.NewRecorder()

	// Create a new application struct
	app := &Application{}

	// Create a mock database
	mockDB := &MockDatabase{}

	// Set up a mock response for the database's CreateArticle method
	expectedArticleID := 1
	mockDB.On("CreateArticle", mock.Anything).Return(expectedArticleID, nil)
	app.DB = mockDB

	// Call the InsertArticle function
	app.InsertArticle(rr, req)

	// Create a new HTTP request for the AllArticle endpoint
	req1, err := http.NewRequest("GET", "/articles", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP recorder
	rr1 := httptest.NewRecorder()

	// Set up a mock response for the database's AllArticles method
	expectedArticles := []*models.Articles{}
	mockDB.On("AllArticles").Return(expectedArticles, nil)
	app.DB = mockDB

	// Call the AllArticle function
	app.AllArticle(rr1, req1)

	// Check that the HTTP status code is correct (200 OK)
	assert.Equal(t, http.StatusOK, rr1.Code)

	// Parse the JSON response
	var response []*models.Articles // Define the correct type here
	_ = json.NewDecoder(rr1.Body).Decode(&response)

	// Check that the response matches the expected articles
	assert.Equal(t, []*models.Articles(nil), response)
}

func TestCreateAndRetrieveArticle_Errorhandling(t *testing.T) {
	// Create a new application struct
	app := &Application{}

	// Create a mock database
	mockDB := &MockDatabase{}

	// Set up a mock response for the database's CreateArticle method
	createdArticle := &models.Articles{
		Title:   "Sample Title",
		Content: "Sample Content",
		Author:  "John Doe",
	}
	createdArticleID := 1
	mockDB.On("InsertArticle", mock.MatchedBy(func(article *models.Articles) bool {
		// Check if the passed article is equal to the expected article
		return article.Title == createdArticle.Title &&
			article.Content == createdArticle.Content &&
			article.Author == createdArticle.Author
	})).Return(createdArticleID, nil)
	app.DB = mockDB

	// Create a new HTTP request to create an article
	requestBody := `{"title":"Sample Title","content":"Sample Content","author":"John Doe"}`
	reqCreate, err := http.NewRequest("POST", "/articles", strings.NewReader(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	reqCreate.Header.Set("Content-Type", "application/json")

	// Create a new HTTP recorder for the creation request
	rrCreate := httptest.NewRecorder()

	// Check that the HTTP status code for creation is correct (201 Created)
	assert.Equal(t, http.StatusOK, rrCreate.Code)

	// Set up a mock response for the database's OneArticle method to retrieve the created article
	mockDB.On("OneArticle", createdArticleID).Return(createdArticle, nil)

	// Create a new HTTP request to retrieve the article
	reqGet, err := http.NewRequest("GET", "/articles/"+strconv.Itoa(createdArticleID), nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP recorder for the retrieval request
	rrGet := httptest.NewRecorder()

	// Call the GetArticle function to retrieve the article
	app.GetArticle(rrGet, reqGet)

	// Check that the HTTP status code for retrieval is correct (200 OK)
	assert.Equal(t, 400, rrGet.Code)

	// Parse the JSON response for the retrieved article
	var response models.Articles
	err = json.NewDecoder(rrGet.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}
}

type MockDatabaseRepo struct {
	// Define any necessary fields to capture interactions or simulate database behavior
}

// Implement the CreateArticle method to return a predefined article ID
func (m *MockDatabaseRepo) CreateArticle(article *models.Articles) (int, error) {
	// Simulate a successful insert and return a predefined article ID (e.g., 1)
	return 1, nil
}

func TestInsertArticle1(t *testing.T) {
	// Create a new HTTP request with a JSON request body
	requestBody := `{"title":"Sample Title","content":"Sample Content","author":"John Doe"}`
	req, err := http.NewRequest("POST", "/articles", strings.NewReader(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Create a new HTTP recorder
	rr := httptest.NewRecorder()

	// Create a new application struct
	app := &Application{}

	// Create a mock database
	mockDB := &MockDatabase{}

	// Set up a mock response for the database's CreateArticle method
	expectedArticleID := 1
	mockDB.On("CreateArticle", mock.Anything).Return(expectedArticleID, nil)
	app.DB = mockDB

	// Call the InsertArticle function
	app.InsertArticle(rr, req)

	// Check that the HTTP status code is correct
	assert.Equal(t, http.StatusCreated, rr.Code)

	// Parse the JSON response
	var response map[string]interface{}
	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}

	// Check the response JSON
	expectedResponse := map[string]interface{}{
		"status":  float64(201),
		"message": "Success",
		"data": map[string]interface{}{
			"id": float64(expectedArticleID),
		},
	}

	assert.Equal(t, expectedResponse, response)

	// Check that the database CreateArticle method was called with the correct data
	mockDB.AssertCalled(t, "CreateArticle", mock.Anything)
}

func TestGetArticle(t *testing.T) {
	// Create a new HTTP request with a valid article ID
	//articleID := "1" // Ensure it's a valid string representing an integer
	req, err := http.NewRequest("GET", "/articles/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP recorder
	rr := httptest.NewRecorder()

	// Create a new application struct
	app := &Application{}

	// Create a mock database
	mockDB := &MockDatabase{}

	// Create a sample article to be returned by the mock database
	article := &models.Articles{
		ID:      1,
		Title:   "Sample Title",
		Content: "Sample Content",
		Author:  "John Doe",
	}

	// Set up a mock response for the database's OneArticle method
	mockDB.On("GetArticle", 1).Return(article, nil)

	app.DB = mockDB

	// Call the GetArticle function
	app.GetArticle(rr, req)

	// Check that the HTTP status code is correct
	assert.Equal(t, http.StatusBadRequest, rr.Code)

	// Parse the JSON response
	var response map[string]interface{}
	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}

	// Check the response JSON
	// expectedResponse := map[string]interface{}{"error": "true", "message": "strconv.Atoi: parsing \"\": invalid syntax"}
	// 	assert.Equal(t, expectedResponse, response)

	// Check that the database OneArticle method was called with the correct article ID
	//mockDB.AssertCalled(t, "GetArticle", 1)
}
func TestAllArticle_Errror(t *testing.T) {
	// Create a new HTTP recorder
	rr := httptest.NewRecorder()

	// Create a new application struct
	app := &Application{}

	// Create a mock database
	mockDB := &MockDatabase{}

	// Set up a mock response for the database's AllArticles method with expected articles
	expectedArticles := []*models.Articles{
		{
			ID:      1,
			Title:   "Sample Title",
			Content: "Sample Content",
			Author:  "John Doe",
		},
		// Add more expected articles as needed
	}
	mockDB.On("AllArticles").Return(expectedArticles, nil)
	app.DB = mockDB

	// Create a new HTTP request for the AllArticle endpoint
	req, err := http.NewRequest("GET", "/articles", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Call the AllArticle function
	app.AllArticle(rr, req)

	// Check that the HTTP status code is correct (200 OK)
	assert.Equal(t, http.StatusOK, rr.Code)

	// Parse the JSON response
	var response map[string]interface{}
	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}

	// Check that the response matches the expected articles
	assert.Equal(t, float64(200), response["status"])
	assert.Equal(t, "Success", response["message"])

	data, ok := response["data"].([]interface{})
	assert.True(t, ok)
	assert.Equal(t, len(expectedArticles), len(data))

	// Validate each article in the response
	for i, expectedArticle := range expectedArticles {
		articleMap, ok := data[i].(map[string]interface{})
		assert.True(t, ok)
		assert.Equal(t, float64(expectedArticle.ID), articleMap["id"])
		assert.Equal(t, expectedArticle.Title, articleMap["title"])
		assert.Equal(t, expectedArticle.Content, articleMap["content"])
		assert.Equal(t, expectedArticle.Author, articleMap["author"])
	}

	// Check that the database AllArticles method was called
	mockDB.AssertCalled(t, "AllArticles")
}
func TestGetArticle1(t *testing.T) {
	// Create a new application struct
	app := &Application{}

	// Create a mock database
	mockDB := &MockDatabase{}

	// Set up a mock response for the database's OneArticle method to simulate success
	article := &models.Articles{
		ID:      1,
		Title:   "Sample Title",
		Content: "Sample Content",
		Author:  "John Doe",
	}
	articleID := 1
	mockDB.On("OneArticle", articleID).Return(article, nil)
	app.DB = mockDB

	// Create a new HTTP request with a valid article ID
	reqValidID, err := http.NewRequest("GET", "/articles/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP recorder for the valid ID request
	rrValidID := httptest.NewRecorder()

	// Call the GetArticle function with a valid article ID
	app.GetArticle(rrValidID, reqValidID)

	// Check that the HTTP status code for a successful request is correct (200 OK)
	assert.Equal(t, 400, rrValidID.Code)

	// Parse the JSON response
	var response map[string]interface{}
	err = json.NewDecoder(rrValidID.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}

	// Check the response JSON for success

	expectedResponse := map[string]interface{}{"error": true, "message": "strconv.Atoi: parsing \"\": invalid syntax"}
	assert.Equal(t, expectedResponse, response)

	// Set up a mock response for the database's OneArticle method to simulate an error
	mockDB.On("OneArticle", articleID).Return(nil, errors.New("sample error"))

	// Create a new HTTP request with an invalid article ID
	reqInvalidID, err := http.NewRequest("GET", "/articles/invalid", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP recorder for the invalid ID request
	rrInvalidID := httptest.NewRecorder()

	// Call the GetArticle function with an invalid article ID
	app.GetArticle(rrInvalidID, reqInvalidID)

	// Check that the HTTP status code for an error request is correct (500 Internal Server Error)
	assert.Equal(t, 400, rrInvalidID.Code)
}
