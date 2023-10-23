package models

// Response
//
// swagger:response Response
type Response struct {
	// Status code of the response
	Status int `json:"status"`
	// Success or error message
	Message string `json:"message"`
	// Any type of Response data or null
	Data interface{} `json:"data"`
}
