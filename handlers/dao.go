package handlers

// Book holds all info related to a book
type Book struct {
	BookID string  `json:"bookId"`
	Name   string  `json:"name"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}
