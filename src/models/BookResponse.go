package models

type BookResponse struct {
    Title  string  `json:"title"`
    Author string  `json:"author"` 
    Price  float64 `json:"price"`
}