package controllers

import (
	"bookStore/db"
	"bookStore/models"
)

type BookController struct {
	database db.DBMS
}

type Controller interface {
	CreateBook(book *models.Book) error
	GetBook(bookID string) (*models.Book, error)
	GetAllBooks() ([]*models.Book, error)
	DeleteBook(bookID string) (int, error)
	UpdateBook(book *models.Book) (int, error)
}

func NewBookController(database db.DBMS) Controller {
	return &BookController{database: database}
}
