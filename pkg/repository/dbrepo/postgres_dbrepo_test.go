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

type DBRepo interface {
	OneArticle(id int) (*models.Articles, error)
}

// Test case using the table driven test
func TestPostgresDBRepo(t *testing.T) {
	tests := []struct {
		name        string
		setupMock   func(mock sqlmock.Sqlmock)
		repoAction  func(repo *PostgresDBRepo) error
		expectedErr error
	}{
		{
			name: "Test CreateTable",
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("CREATE TABLE IF NOT EXISTS articles").
					WillReturnResult(sqlmock.NewResult(0, 0))
			},
			repoAction: func(repo *PostgresDBRepo) error {
				repo.CreateTable()
				return nil
			},
			expectedErr: nil,
		},
		{
			name: "Test AllArticles",
			setupMock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "title", "content", "author"}).
					AddRow(1, "Title1", "Content1", "Author1").
					AddRow(2, "Title2", "Content2", "Author2")

				mock.ExpectQuery("SELECT id, title, content, author FROM articles").
					WillReturnRows(rows)
			},
			repoAction: func(repo *PostgresDBRepo) error {
				_, err := repo.AllArticles()
				return err
			},
			expectedErr: nil,
		},
		{
			name: "Test OneArticle (article found)",
			setupMock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "title", "content", "author"}).
					AddRow(1, "Title1", "Content1", "Author1")

				mock.ExpectQuery("SELECT id, title, content, author FROM articles WHERE id = \\$1").
					WithArgs(1).
					WillReturnRows(rows)
			},
			repoAction: func(repo *PostgresDBRepo) error {
				_, err := repo.OneArticle(1)
				return err
			},
			expectedErr: nil,
		},
		{
			name: "Test OneArticle (article not found)",
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT id, title, content, author FROM articles WHERE id = \\$1").
					WithArgs(2).
					WillReturnError(sql.ErrNoRows)
			},
			repoAction: func(repo *PostgresDBRepo) error {
				_, err := repo.OneArticle(2)
				return err
			},
			//expectedErr: models.ErrNoRecord,
		},
		{
			name: "Test CreateArticle",
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("INSERT INTO articles").
					WithArgs("Title1", "Content1", "Author1").
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
			},
			repoAction: func(repo *PostgresDBRepo) error {
				_, err := repo.CreateArticle(&models.Articles{
					Title:   "Title1",
					Content: "Content1",
					Author:  "Author1",
				})
				return err
			},
			expectedErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("Failed to create mock database: %v", err)
			}
			defer db.Close()

			repo := &PostgresDBRepo{DB: db}
			test.setupMock(mock)
			err = test.repoAction(repo)

			if err != test.expectedErr {
				t.Errorf("Expected error: %v, got: %v", test.expectedErr, err)
			}
		})
	}
}

func TestOneArticle_error(t *testing.T) {
	// Create a new mock DB connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock DB: %v", err)
	}
	defer db.Close()

	// Create a PostgresDBRepo with the mock database connection.
	repo := &PostgresDBRepo{DB: db}

	// Define the expected SQL query
	query := "SELECT id, title, content, author FROM articles WHERE id = ?"

	// Expect the SQL query with id = 1 to return sql.ErrNoRows
	mock.ExpectQuery(query).
		WithArgs(1).
		WillReturnError(sql.ErrNoRows)

	// Test the OneArticle function
	fetchedArticle, _ := repo.OneArticle(1)

	// Verify that an error of sql.ErrNoRows is returned, indicating no article found
	assert.Nil(t, fetchedArticle, "Expected no article to be found, but got %v", fetchedArticle)

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
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
