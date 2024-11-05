package models

type Book struct {
    ID       string  `json:"id" bson:"_id,omitempty"`
    Title    string  `json:"title" bson:"title"`
    Author   string  `json:"author" bson:"author"`
    Year     int     `json:"year" bson:"year"`
    Stock    int     `json:"stock" bson:"stock"`
    Price    float64 `json:"price" bson:"price"`
}