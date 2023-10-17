package api

import (
	"backend/pkg/models"
	"backend/pkg/utility"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

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

	// Prepare the response map
	response := map[string]interface{}{
		"status":  200,
		"message": "Success",
		"data":    []map[string]interface{}{},
	}

	// Check if there are articles, then add them to the response map
	if len(articles) > 0 {
		for _, article := range articles {

			response["data"] = append(response["data"].([]map[string]interface{}), map[string]interface{}{"id": article.ID, "title": article.Title, "content": article.Content, "author": article.Author})
		}
	}

	// Set the response headers and write the JSON response
	utility.WriteJSON(w, http.StatusOK, response)
}

func (app *Application) GetArticle(w http.ResponseWriter, r *http.Request) {
	// Get the article ID from the URL parameter
	id := chi.URLParam(r, "id")
	articleID, err := strconv.Atoi(id)
	if err != nil {
		utility.ErrorJSON(w, err)
		return
	}

	// Retrieve the article from the database
	article, _ := app.DB.OneArticle(articleID)

	// Prepare the response map
	response := map[string]interface{}{
		"status":  200,
		"message": "Success",
		"data":    []map[string]interface{}{},
	}

	// Check if the article is not nil, then add it to the response map

	response["data"] = append(response["data"].([]map[string]interface{}), map[string]interface{}{"id": article.ID, "title": article.Title, "content": article.Content, "author": article.Author})

	// Set the response headers and write the JSON response
	utility.WriteJSON(w, http.StatusOK, response)
}

func (app *Application) InsertArticle(w http.ResponseWriter, r *http.Request) {
	// Parse the JSON request body into an Article struct
	var article models.Articles
	_ = utility.ReadJSON(w, r, &article)

	// Insert the article into the database
	articleID, _ := app.DB.CreateArticle(&article)

	// Prepare the response JSON
	response := map[string]interface{}{
		"status":  201,
		"message": "Success",
		"data": map[string]interface{}{
			"id": articleID,
		},
	}

	// Set the response headers and write the JSON response
	utility.WriteJSON(w, http.StatusCreated, response)
}
