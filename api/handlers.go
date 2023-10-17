package api

import (
	"backend/pkg/models"
	"backend/pkg/utility"
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// cd api then fire command to create mockgen interface -->> mockgen -source=handlers.go -destination=mocks/mock_handlers.go -package=mocks
type DBInterface interface {
	Connection() *sql.DB
	CreateTable()
	AllArticles() ([]*models.Articles, error)
	CreateArticle(article *models.Articles) (int, error)
	OneArticle(id int) (*models.Articles, error)
}

type UtilityInterface interface {
	WriteJSON(w http.ResponseWriter, status int, data interface{}) error
	ReadJSON(w http.ResponseWriter, r *http.Request, data interface{}) error
}

// Home displays the status of the api, as JSON.
func (app *Application) Home(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Go articles up and running",
		Version: "1.0.0",
	}

	_ = utility.WriteJSON(w, http.StatusOK, payload)
}

func (app *Application) AllArticle(w http.ResponseWriter, r *http.Request) {
	// Retrieve the list of articles from the database
	articles, _ := app.DB.AllArticles()

	// Create the response struct
	var response models.Response

	response.Status = http.StatusOK
	response.Message = "Success"

	var dataList []models.Articles // Slice to hold the articles

	// Check if there are articles, then add them to the response
	for _, article := range articles {
		articleData := models.Articles{
			ID:      article.ID,
			Title:   article.Title,
			Content: article.Content,
			Author:  article.Author,
		}
		dataList = append(dataList, articleData)
	}

	// Set the response data as a slice of articles
	response.Data = dataList
	if len(dataList) <= 1 {
		log.Println("No article found for the given ID")
		utility.WriteJSON(w, http.StatusNotFound, models.Response{Data: nil, Status: http.StatusNotFound, Message: "No article found"})
		return
	}
	// Set the response headers and write the JSON response
	utility.WriteJSON(w, http.StatusOK, response)
}

func (app *Application) GetArticle(w http.ResponseWriter, r *http.Request) {
	// Get the article ID from the URL parameter
	id := chi.URLParam(r, "id")
	articleID, _ := strconv.Atoi(id)

	// Retrieve the article from the database
	article, _ := app.DB.OneArticle(articleID)

	// Create the response struct
	var response models.Response
	response.Status = http.StatusOK
	response.Message = "Success"

	if article != nil {
		response.Data = models.Articles{
			ID:      article.ID,
			Title:   article.Title,
			Content: article.Content,
			Author:  article.Author,
		}
	}
	if article == nil {
		log.Println("No article found for the ID")
		utility.WriteJSON(w, http.StatusNotFound, models.Response{Data: nil, Status: http.StatusNotFound, Message: "No article found for the ID"})
		return
	}
	utility.WriteJSON(w, http.StatusOK, response)

}

func (app *Application) InsertArticle(w http.ResponseWriter, r *http.Request) {
	// Parse the JSON request body into an Article struct
	var article models.Articles
	_ = utility.ReadJSON(w, r, &article)

	// Insert the article into the database
	articleID, err := app.DB.CreateArticle(&article)
	if err != nil {
		// Handle the error here
		log.Println("Article not created due to database error")
		utility.WriteJSON(w, http.StatusInternalServerError, models.Response{Data: nil, Status: http.StatusInternalServerError, Message: "Article not created due to database error"})
		return
	}
	// Create the response struct
	var response models.Response
	response.Status = http.StatusCreated
	response.Message = "Success"

	// Prepare the response JSON
	response.Data = models.Articles{
		ID: articleID,
	}

	// Set the response headers and write the JSON response
	utility.WriteJSON(w, http.StatusCreated, response)
}
