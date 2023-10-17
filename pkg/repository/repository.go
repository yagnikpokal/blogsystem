package repository

import (
	"backend/pkg/models"
	"database/sql"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	CreateTable()
	AllArticles() ([]*models.Articles, error)
	CreateArticle(article *models.Articles) (int, error)
	OneArticle(id int) (*models.Articles, error)
}
