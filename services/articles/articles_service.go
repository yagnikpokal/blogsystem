package services

import (
	"backend/pkg/models"
	"backend/pkg/repository/dbrepo"
)

type ArticleServices interface {
	GetAllArticles() ([]models.Articles, error)
	GetArticleByID(id int) (*models.Articles, error)
	CreateArticle(article *models.Articles) (int, error)
}

type ArticleService struct {
	repo dbrepo.DatabaseRepo
}

func NewArticleService(repo dbrepo.DatabaseRepo) *ArticleService {
	return &ArticleService{
		repo: repo,
	}
}

func (s *ArticleService) GetAllArticles() ([]models.Articles, error) {
	// Call the GetAllArticles method from the repository
	return s.repo.AllArticles()
}

func (s *ArticleService) GetArticleByID(id int) (*models.Articles, error) {
	return s.repo.OneArticle(id)
}

func (s *ArticleService) CreateArticle(article *models.Articles) (int, error) {
	return s.repo.CreateArticle(article)
}
