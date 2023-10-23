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
		Data    int    `json:"data"`
	}
}

// IDParameter represents the 'id' parameter in the Swagger schema.
//
// swagger:parameters idParameter
type IDParameter struct {
	// in: path
	// required: true
	// type: int
	// example: 1
	ID int `json:"id"`
}
