package dbrepo

import (
	"backend/pkg/models"
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type PostgresDBRepo struct {
	DB *sql.DB
}

const dbTimeout = time.Second * 3

// Connection returns underlying connection pool.
func (m *PostgresDBRepo) Connection() *sql.DB {
	return m.DB
}

// Create a table if not exist
func (m *PostgresDBRepo) CreateTable() {

	createTableSQL := `
        CREATE TABLE IF NOT EXISTS articles (
            id SERIAL PRIMARY KEY,
            title TEXT NOT NULL,
            content TEXT NOT NULL,
            author TEXT NOT NULL
        );`
	_, err := m.DB.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Table 'articles' created successfully!")
}

// Return all articles
func (m *PostgresDBRepo) AllArticles() ([]*models.Articles, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
        SELECT
            id, title, content, author
        FROM
            articles
        ORDER BY
            title
    `

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		log.Println("Error in getting query: ", err)
		return nil, err
	}
	defer rows.Close()

	var articlesList []*models.Articles

	for rows.Next() {
		var article models.Articles
		err := rows.Scan(
			&article.ID,
			&article.Title,
			&article.Content,
			&article.Author,
		)
		if err != nil {
			log.Println("Error in getting next row: ", err)
			return nil, err
		}

		articlesList = append(articlesList, &article)
	}

	return articlesList, nil
}

// Retrive one article
func (m *PostgresDBRepo) OneArticle(id int) (*models.Articles, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
        SELECT
            id, title, content, author
        FROM
            articles
        WHERE
            id = $1
    `

	var article models.Articles
	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&article.ID,
		&article.Title,
		&article.Content,
		&article.Author,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No article found: ", err)

			return nil, nil // Article not found
		}
		log.Println("Error in getting query: ", err)
		return nil, err // Other error
	}

	return &article, nil
}

// Create new article
func (m *PostgresDBRepo) CreateArticle(article *models.Articles) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
        INSERT INTO articles (title, content, author)
        VALUES ($1, $2, $3)
        RETURNING id
    `

	var articleID int
	err := m.DB.QueryRowContext(ctx, query, article.Title, article.Content, article.Author).Scan(&articleID)
	if err != nil {
		log.Println("Error in getting query: ", err)
		return 0, err
	}

	return articleID, nil
}
