package dbrepo

import (
	"backend/pkg/models"
	"database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/jackc/pgx/v4/stdlib" // Import the PostgreSQL driver
	"github.com/stretchr/testify/assert"
)

func TestCreateArticle(t *testing.T) {
	// Create a new mock DB connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock DB: %v", err)
	}
	defer db.Close()

	// Create a PostgresDBRepo with the mock database connection.
	repo := &PostgresDBRepo{DB: db}

	// Create a sample article.
	article := &models.Articles{
		Title:   "Test Article",
		Content: "Test Content",
		Author:  "Test Author",
	}

	// Expect the INSERT statement and define its expected result.
	mock.ExpectQuery("INSERT INTO articles (.+) VALUES (.+) RETURNING id").
		WithArgs(article.Title, article.Content, article.Author).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	// Test the CreateArticle function.
	articleID, err := repo.CreateArticle(article)
	if err != nil {
		t.Fatalf("Error creating article: %v", err)
	}

	// Add assertions to verify the result.
	if articleID != 1 {
		t.Errorf("Expected article ID 1, but got %d", articleID)
	}

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestOneArticle1(t *testing.T) {
	// Create a new mock DB connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock DB: %v", err)
	}
	defer db.Close()

	// Create a PostgresDBRepo with the mock database connection.
	repo := &PostgresDBRepo{DB: db}

	// Create a sample article in the mock database.
	article := &models.Articles{
		Title:   "Test Article",
		Content: "Test Content",
		Author:  "Test Author",
	}

	// Define the expected SQL query
	query := "SELECT id, title, content, author FROM articles WHERE id = ?"

	// Set the expectation for the SQL query
	mock.ExpectQuery(query).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "content", "author"}).
			AddRow(1, article.Title, article.Content, article.Author))

	// Test the OneArticle function
	fetchedArticle, err := repo.OneArticle(1)
	if err != nil {
		t.Fatalf("Error fetching article: %v", err)
	}

	// Use assertions to verify the result
	assert.NotNil(t, fetchedArticle, "Expected an article, but got nil")
	assert.Equal(t, 1, fetchedArticle.ID, "Expected article with ID 1, but got ID %d", fetchedArticle.ID)

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestAllArticles(t *testing.T) {
	// Create a new mock DB connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock DB: %v", err)
	}
	defer db.Close()

	// Create a PostgresDBRepo with the mock database connection.
	repo := &PostgresDBRepo{DB: db}

	// Define the expected SQL query and result.
	query := "SELECT id, title, content, author FROM articles ORDER BY title"
	expectedRows := sqlmock.NewRows([]string{"id", "title", "content", "author"}).
		AddRow(1, "Article 1", "Content 1", "Author 1").
		AddRow(2, "Article 2", "Content 2", "Author 2")

	// Set the expectation for the SQL query.
	mock.ExpectQuery(query).WillReturnRows(expectedRows)

	// Test the AllArticles function.
	articlesList, err := repo.AllArticles()
	if err != nil {
		t.Fatalf("Error fetching articles: %v", err)
	}

	// Add assertions to verify the result.
	if len(articlesList) != 2 {
		t.Errorf("Expected 2 articles, but got %d", len(articlesList))
	}

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}
func TestOneArticle(t *testing.T) {
	// Create a new mock DB connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock DB: %v", err)
	}
	defer db.Close()

	// Create a PostgresDBRepo with the mock database connection.
	repo := &PostgresDBRepo{DB: db}

	// Create a sample article in the mock database.
	article := &models.Articles{
		Title:   "Test Article",
		Content: "Test Content",
		Author:  "Test Author",
	}

	// Define the expected SQL query
	query := "SELECT id, title, content, author FROM articles WHERE id = ?"

	// Set the expectation for the SQL query
	mock.ExpectQuery(query).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "content", "author"}).
			AddRow(1, article.Title, article.Content, article.Author))

	// Test the OneArticle function
	fetchedArticle, err := repo.OneArticle(1)
	if err != nil {
		t.Fatalf("Error fetching article: %v", err)
	}

	// Use assertions to verify the result
	assert.NotNil(t, fetchedArticle, "Expected an article, but got nil")
	assert.Equal(t, 1, fetchedArticle.ID, "Expected article with ID 1, but got ID %d", fetchedArticle.ID)

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}
func TestConnection(t *testing.T) {
	// Replace with your actual database connection parameters
	dsn := "user=username dbname=mydb sslmode=disable"
	testDB, err := sql.Open("pgx", dsn) // Use "pgx" as the driver name
	if err != nil {
		t.Fatalf("Error creating test DB: %v", err)
	}
	defer testDB.Close()

	repo := &PostgresDBRepo{DB: testDB}

	connection := repo.Connection()
	assert.Equal(t, testDB, connection, "Connection method did not return the expected *sql.DB instance.")
}
func TestConnections(t *testing.T) {
	// Create a new mock DB connection
	db, _, _ := sqlmock.New()
	repo := &PostgresDBRepo{DB: db}

	// Call the Connection method
	connection := repo.Connection()

	// Check if the returned connection is the same as the mock DB
	assert.Equal(t, db, connection, "Connection method did not return the expected *sql.DB instance.")
}

func TestCreateTable(t *testing.T) {
	// Create a new mock DB connection
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := &PostgresDBRepo{DB: db}

	// Expect the CREATE TABLE statement
	mock.ExpectExec("CREATE TABLE IF NOT EXISTS articles").
		WillReturnResult(sqlmock.NewResult(0, 0))

	// Call the CreateTable method
	repo.CreateTable()

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}
func TestCreateArticleError(t *testing.T) {
	// Create a new mock DB connection
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := &PostgresDBRepo{DB: db}

	// Expect the INSERT statement to return an error
	mock.ExpectQuery("INSERT INTO articles").
		WillReturnError(fmt.Errorf("Test error"))

	// Create a sample article
	article := &models.Articles{
		Title:   "Test Article",
		Content: "Test Content",
		Author:  "Test Author",
	}

	// Call CreateArticle, which should return an error
	articleID, err := repo.CreateArticle(article)

	// Check if the returned error is as expected
	assert.Error(t, err, "Expected an error")
	assert.Contains(t, err.Error(), "Test error", "Error message not as expected")

	// Ensure that articleID is 0
	assert.Equal(t, 0, articleID, "Expected articleID to be 0")
}
func TestAllArticlesError(t *testing.T) {
	// Create a new mock DB connection
	db, mock, _ := sqlmock.New()
	defer db.Close()

	repo := &PostgresDBRepo{DB: db}

	// Expect the SELECT statement to return an error
	mock.ExpectQuery("SELECT (.+) FROM articles").
		WillReturnError(fmt.Errorf("Test query error"))

	// Call AllArticles, which should return an error
	articles, err := repo.AllArticles()

	// Check if the returned error is as expected
	assert.Error(t, err, "Expected an error")
	assert.Contains(t, err.Error(), "Test query error", "Error message not as expected")

	// Ensure that articles is nil
	assert.Nil(t, articles, "Expected articles to be nil")

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}
