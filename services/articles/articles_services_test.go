package services

import (
	"errors"
	"testing"

	"backend/mocks" // Import the generated mock package
	appconst "backend/pkg/appconstant"
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

	// Define the test cases
	testCases := []struct {
		description  string
		expectedData []models.Article
		expectedErr  error
		mockFunc     func() ([]models.Article, error)
	}{
		{
			description:  "Successful case",
			expectedData: []models.Article{{ID: 1, Title: "Article 1", Content: "Content 1"}, {ID: 2, Title: "Article 2", Content: "Content 2"}},
			expectedErr:  nil,
			mockFunc: func() ([]models.Article, error) {
				return []models.Article{{ID: 1, Title: "Article 1", Content: "Content 1"}, {ID: 2, Title: "Article 2", Content: "Content 2"}}, nil
			},
		},
		{
			description:  "Negative test case",
			expectedData: nil,
			expectedErr:  errors.New(appconst.Noarticlefound),
			mockFunc: func() ([]models.Article, error) {
				return nil, errors.New(appconst.Noarticlefound)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			// Set up expectations for the mock repo
			mockDB.EXPECT().AllArticles().DoAndReturn(testCase.mockFunc)

			// Call the GetAllArticles method
			articles, err := service.GetAllArticles()

			// Check the result
			assert.Equal(t, testCase.expectedErr, err)
			assert.Equal(t, testCase.expectedData, articles)
		})
	}
}

func TestArticleService_CreateArticle(t *testing.T) {
	// Create a new instance of the mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDB := mocks.NewMockDBInterface(ctrl)

	// Create a new instance of the ArticleService with the mock repo
	service := NewArticleService(mockDB)

	// Define the test cases
	testCases := []struct {
		description       string
		articleToCreate   *models.Article
		expectedArticleID int
		expectedErr       error
		mockFunc          func(article *models.Article) (int, error)
	}{
		{
			description:       "Successful creation",
			articleToCreate:   &models.Article{Title: "New Article", Content: "New Content"},
			expectedArticleID: 1,
			expectedErr:       nil,
			mockFunc: func(article *models.Article) (int, error) {
				return 1, nil
			},
		},
		{
			description:       "Negative test case",
			articleToCreate:   &models.Article{Title: "New Article", Content: "New Content"},
			expectedArticleID: 0,
			expectedErr:       errors.New(appconst.Articlenotcreated),
			mockFunc: func(article *models.Article) (int, error) {
				return 0, errors.New(appconst.Articlenotcreated)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			// Set up expectations for the mock repo
			mockDB.EXPECT().CreateArticle(testCase.articleToCreate).DoAndReturn(testCase.mockFunc)

			// Call the CreateArticle method
			createdArticleID, err := service.CreateArticle(testCase.articleToCreate)

			// Check the result
			assert.Equal(t, testCase.expectedErr, err)
			assert.Equal(t, testCase.expectedArticleID, createdArticleID)
		})
	}
}

func TestArticleService_GetArticleByID(t *testing.T) {
	// Create a new instance of the mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockDB := mocks.NewMockDBInterface(ctrl)

	// Create a new instance of the ArticleService with the mock repo
	service := NewArticleService(mockDB)

	// Define the test cases
	testCases := []struct {
		description     string
		articleID       int
		expectedArticle *models.Article
		expectedErr     error
		mockFunc        func(id int) (*models.Article, error)
	}{
		{
			description:     "Successful case",
			articleID:       1,
			expectedArticle: &models.Article{ID: 1, Title: "Article 1", Content: "Content 1"},
			expectedErr:     nil,
			mockFunc: func(id int) (*models.Article, error) {
				return &models.Article{ID: 1, Title: "Article 1", Content: "Content 1"}, nil
			},
		},
		{
			description:     "Negative test case",
			articleID:       2,
			expectedArticle: nil,
			expectedErr:     errors.New(appconst.NoArticleforid),
			mockFunc: func(id int) (*models.Article, error) {
				return nil, errors.New(appconst.NoArticleforid)
			},
		},
		{
			description:     "Article not found",
			articleID:       3,
			expectedArticle: nil,
			expectedErr:     errors.New(appconst.Noarticlefound),
			mockFunc: func(id int) (*models.Article, error) {
				return nil, errors.New(appconst.Noarticlefound)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			// Set up expectations for the mock repo
			mockDB.EXPECT().OneArticle(testCase.articleID).DoAndReturn(testCase.mockFunc)

			// Call the GetArticleByID method
			article, err := service.GetArticleByID(testCase.articleID)

			// Check the result
			assert.Equal(t, testCase.expectedErr, err)
			assert.Equal(t, testCase.expectedArticle, article)
		})
	}
}
