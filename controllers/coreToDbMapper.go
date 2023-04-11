package controllers

import (
	"bookStore/db"
	"bookStore/models"
)

func toDBReq(req *models.Book) *db.Book {
	return &db.Book{
		BookID: req.BookID,
		Name:   req.Name,
		Author: req.Author,
		Price:  req.Price,
	}
}

func toCoreControllerBook(req *db.Book) *models.Book {
	return &models.Book{
		BookID: req.BookID,
		Name:   req.Name,
		Author: req.Author,
		Price:  req.Price,
	}
}

func toCoreControllerBooks(req []*db.Book) []*models.Book {
	var resp []*models.Book
	for _, book := range req {
		resp = append(resp, &models.Book{
			BookID: book.BookID,
			Name:   book.Name,
			Author: book.Author,
			Price:  book.Price,
		})
	}
	return resp
}
