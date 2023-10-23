package api

import (
	appconst "backend/pkg/appconstant"
	"backend/pkg/models"
	"backend/pkg/utility"
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

/*
Do not remove the following interfaces, as they are used for mock data generation.
Command to generate mock data: make mocks
*/
type DBInterface interface {
	Connection() *sql.DB
	CreateTable()
	AllArticles() ([]models.Article, error)
	CreateArticle(article *models.Article) (int, error)
	OneArticle(id int) (*models.Article, error)
}

type UtilityInterface interface {
	WriteJSON(w http.ResponseWriter, status int, data interface{}) error
	ReadJSON(w http.ResponseWriter, r *http.Request, data interface{}) error
}

// HealthCheck performs a basic health check of the service.
//
// swagger:route GET / healthCheck
//
// Performs a basic health check of the service.
//
// Responses:
//   200: SuccessResponse
//	500: ErrorResponse

func (app *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	// Create the response struct
	var response models.Response

	response.Status = http.StatusOK
	response.Message = appconst.Success
	response.Data = appconst.Serverup

	utility.WriteJSON(w, http.StatusOK, response)
}

// AllArticle retrieves a list of articles.
//
// swagger:route GET /articles allArticle
//
// Retrieve a list of articles.
//
// Responses:
//
//	200: ArticleListResponse
//	500: ErrorResponse
func (app *Application) AllArticle(w http.ResponseWriter, r *http.Request) {
	// Retrieve the list of articles from the database
	articles, err := app.ArticleService.GetAllArticles()
	if err != nil {
		// Handle the error
		log.Println(appconst.Errorconst, err)
		utility.WriteJSON(w, http.StatusInternalServerError, models.Response{Data: nil, Status: http.StatusInternalServerError, Message: appconst.Errorconst + err.Error()})
		return
	}
	// Create the response struct
	var response models.Response

	response.Status = http.StatusOK
	response.Message = appconst.Success
	// Set the response data as a slice of articles
	response.Data = articles

	// Set the response headers and write the JSON response
	utility.WriteJSON(w, http.StatusOK, response)
}

// swagger:route GET /articles/{id} idParameter
//
// Retrieve an article by its ID.
// ---
// produces:
// - application/json
// Responses:
//
//	200: ArticleListResponse
//	500: ErrorResponse

func (app *Application) GetArticle(w http.ResponseWriter, r *http.Request) {
	// Get the article ID from the URL parameter
	id := chi.URLParam(r, "id")
	articleID, err := strconv.Atoi(id)
	if err != nil {
		log.Println(appconst.Parsingarticle, err)
		utility.WriteJSON(w, http.StatusBadRequest, models.Response{Data: nil, Status: http.StatusBadRequest, Message: appconst.Parsingarticle + err.Error()})
		return
	}
	// Retrieve the article from the service
	article, err := app.ArticleService.GetArticleByID(articleID)
	if err != nil {
		// Handle the error
		log.Println(appconst.Retrivearticle, err)
		utility.WriteJSON(w, http.StatusInternalServerError, models.Response{Data: nil, Status: http.StatusInternalServerError, Message: appconst.Retrivearticle + err.Error()})
		return
	}

	var response models.Response
	response.Status = http.StatusOK
	response.Message = appconst.Success
	response.Data = article
	utility.WriteJSON(w, http.StatusOK, response)
}

// swagger:operation POST /articles InsertArticle
// ---
// summary: Create an article.
// description: Parses a JSON request to create a new article and returns the result.
// parameters:
// - name: article
//   in: body
//   description: The article data to be created.
//   required: true
//   schema:
//     $ref: '#/definitions/Article'
// responses:
//   201:
//     description: Created
//     $ref: '#/responses/ArticleResponse'
//   500:
//     $ref: '#/responses/ErrorResponse'

func (app *Application) InsertArticle(w http.ResponseWriter, r *http.Request) {
	// Parse the JSON request body into an Article struct
	var article models.Article
	err := utility.ReadJSON(w, r, &article)
	if err != nil {
		log.Println(appconst.JSONparsing, err)
		utility.WriteJSON(w, http.StatusBadRequest, models.Response{Data: nil, Status: http.StatusBadRequest, Message: appconst.JSONparsing})
		return
	}

	// Insert the article into the service
	articleID, err := app.ArticleService.CreateArticle(&article)
	if err != nil {
		// Handle the error here
		log.Println(appconst.Articlenotcreated, err)
		utility.WriteJSON(w, http.StatusInternalServerError, models.Response{Data: nil, Status: http.StatusInternalServerError, Message: appconst.Articlenotcreated + err.Error()})
		return
	}

	// Create the response struct
	var response models.Response
	response.Status = http.StatusCreated
	response.Message = appconst.Success

	// Prepare the response JSON
	response.Data = models.Article{
		ID: articleID,
	}

	// Set the response headers and write the JSON response
	utility.WriteJSON(w, http.StatusCreated, response)
}
