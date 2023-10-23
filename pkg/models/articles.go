package models

// Article
//
// swagger:response Article
type Article struct {
	// ID of the article
	// in: int
	ID int `json:"id"`
	// Title of the article
	// in: string
	Title string `json:"title,omitempty"`
	// Content of the article
	// in: string
	Content string `json:"content,omitempty"`
	// Author of the article
	// in: string
	Author string `json:"author,omitempty"`
}
