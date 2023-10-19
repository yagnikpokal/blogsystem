package dbrepo

import (
	appconst "backend/pkg/appconstant"
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
type DatabaseRepo interface {
	Connection() *sql.DB
	CreateTable()
	AllArticles() ([]models.Article, error)
	CreateArticle(article *models.Article) (int, error)
	OneArticle(id int) (*models.Article, error)
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

	fmt.Println(appconst.CreateArticleTable)
}

// Return all articles
func (m *PostgresDBRepo) AllArticles() ([]models.Article, error) {
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
		log.Println(appconst.Queryerror, err)
		return nil, err
	}
	defer rows.Close()

	var articlesList []models.Article

	for rows.Next() {
		var article models.Article
		err := rows.Scan(
			&article.ID,
			&article.Title,
			&article.Content,
			&article.Author,
		)
		if err != nil {
			log.Println(appconst.Nextrow, err)
			return nil, err
		}

		articlesList = append(articlesList, article)
	}

	return articlesList, nil
}

// Retrive one article
func (m *PostgresDBRepo) OneArticle(id int) (*models.Article, error) {
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

	var article models.Article
	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&article.ID,
		&article.Title,
		&article.Content,
		&article.Author,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(appconst.NoArticleforid, err)
			return nil, err // Article not found
		}
		log.Println(appconst.Queryerror, err)
		return nil, err // Other error
	}

	return &article, nil
}

// Create new article
func (m *PostgresDBRepo) CreateArticle(article *models.Article) (int, error) {
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
		log.Println(appconst.Queryerror, err)
		return 0, err
	}

	return articleID, nil
}
