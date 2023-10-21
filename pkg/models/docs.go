package models

// ArticleListResponse
//
// swagger:response ArticleListResponse
type ArticleListResponse struct {
	// in: body
	Body struct {
		Status  int       `json:"status"`
		Message string    `json:"message"`
		Data    []Article `json:"data"`
	}
}

// ArticleResponse
//
// swagger:response ArticleResponse
type ArticleResponse struct {
	// in: body
	Body struct {
		Status  int     `json:"status"`
		Message string  `json:"message"`
		Data    Article `json:"data"`
	}
}

// SuccessResponse
//
// swagger:response SuccessResponse
type SuccessResponse struct {
	// in: body
	Body Response
}

// ErrorResponse
//
// swagger:response ErrorResponse
type ErrorResponse struct {
	// in: body
	Body Response
}

// ResponseCreateArticle
// swagger:response ResponseCreateArticle
type ResponseCreateArticle struct {
	// in: body
	Body struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Data    ID     `json:"data"`
	}
}

type ID struct {
	// ID of the article
	// in: int
	ID int `json:"id"`
}
