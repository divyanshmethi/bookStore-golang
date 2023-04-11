package models

//Todo: Add publisher functionality later
// Publisher holds all info related to a publisher
//type Publisher struct {
//	ID            string `json:"id"`
//	Name          string `json:"name"`
//	Address       string `json:"address"`
//	Email         string `json:"email"`
//	ContactNumber string `json:"contactNumber"`
//}

// Book holds all info related to a book
type Book struct {
	BookID string `json:"bookId"`
	Name   string `json:"name"`
	Author string `json:"author"`
	//Publisher *Publisher `json:"publisher"`
	Price float64 `json:"price"`
}
