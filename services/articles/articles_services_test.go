package services

import (
	"testing"

	"backend/mocks" // Import the generated mock package
	"backend/pkg/models"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestArticleService_GetAllArticles(t *testing.T) {
	// Create a new instance of the mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDB := mocks.NewMockDBInterface(ctrl)

	// Create a new instance of the ArticleService with the mock repo
	service := NewArticleService(mockDB)

	// Define expected data for the GetAllArticles method
	expectedData := []models.Articles{
		{ID: 1, Title: "Article 1", Content: "Content 1"},
		{ID: 2, Title: "Article 2", Content: "Content 2"},
	}

	// Define the expected error (nil in this case)

	// Set up expectations for the mock repo
	mockDB.EXPECT().AllArticles().Return(expectedData, nil)

	// Call the GetAllArticles method
	articles, err := service.GetAllArticles()

	// Check the result
	assert.Nil(t, err) // Expect no error
	assert.Equal(t, expectedData, articles)
}
func TestArticleService_GetArticleByID(t *testing.T) {
	// Create a new instance of the mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDB := mocks.NewMockDBInterface(ctrl)

	// Create a new instance of the ArticleService with the mock repo
	service := NewArticleService(mockDB)

	// Define the ID for the article to retrieve
	articleID := 1

	// Define the expected article data
	expectedArticle := &models.Articles{
		ID:      1,
		Title:   "Article 1",
		Content: "Content 1",
	}

	// Set up expectations for the mock repo
	mockDB.EXPECT().OneArticle(articleID).Return(expectedArticle, nil)

	// Call the GetArticleByID method
	article, err := service.GetArticleByID(articleID)

	// Check the result
	assert.Nil(t, err) // Expect no error
	assert.Equal(t, expectedArticle, article)
}

func TestArticleService_CreateArticle(t *testing.T) {
	// Create a new instance of the mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDB := mocks.NewMockDBInterface(ctrl)

	// Create a new instance of the ArticleService with the mock repo
	service := NewArticleService(mockDB)

	// Define the article to create
	articleToCreate := &models.Articles{
		Title:   "New Article",
		Content: "New Content",
	}

	// Define the expected article ID upon creation
	expectedArticleID := 1

	// Set up expectations for the mock repo
	mockDB.EXPECT().CreateArticle(articleToCreate).Return(expectedArticleID, nil)

	// Call the CreateArticle method
	createdArticleID, err := service.CreateArticle(articleToCreate)

	// Check the result
	assert.Nil(t, err) // Expect no error
	assert.Equal(t, expectedArticleID, createdArticleID)
}
