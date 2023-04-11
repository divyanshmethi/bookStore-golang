package controllers

import (
	"bookStore/models"
)

func (impl *BookController) CreateBook(book *models.Book) error {
	dbReq := toDBReq(book)
	return impl.database.AddBook(dbReq)
}

func (impl *BookController) GetBook(bookID string) (*models.Book, error) {
	res, err := impl.database.GetBook(bookID)
	if err != nil || res == nil {
		return nil, err
	}
	return toCoreControllerBook(res), err
}

func (impl *BookController) GetAllBooks() ([]*models.Book, error) {
	res, err := impl.database.GetAllBooks()
	if err != nil || res == nil {
		return nil, err
	}
	return toCoreControllerBooks(res), err
}

func (impl *BookController) DeleteBook(bookID string) (int, error) {
	return impl.database.DeleteBook(bookID)
}

func (impl *BookController) UpdateBook(book *models.Book) (int, error) {
	dbReq := toDBReq(book)
	return impl.database.UpdateBook(dbReq)
}
