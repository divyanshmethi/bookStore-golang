package controllers

import (
	"bookStore/models"
)

func (impl *BookController) CreateBook(book *models.Book) error {
	return impl.database.AddBook(book)
}

func (impl *BookController) GetBook(bookID string) (*models.Book, error) {
	return impl.database.GetBook(bookID)
}

func (impl *BookController) GetAllBooks() ([]*models.Book, error) {
	return impl.database.GetAllBooks()
}

func (impl *BookController) DeleteBook(bookID string) (int, error) {
	return impl.database.DeleteBook(bookID)
}

func (impl *BookController) UpdateBook(book *models.Book) (int, error) {
	return impl.database.UpdateBook(book)
}
