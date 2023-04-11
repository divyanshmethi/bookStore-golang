package handlers

import "bookStore/models"

func toControllerReq(req *Book) *models.Book {
	return &models.Book{
		BookID: req.BookID,
		Name:   req.Name,
		Author: req.Author,
		Price:  req.Price,
	}
}
