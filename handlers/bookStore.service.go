package handlers

import (
	"bookStore/controllers"
)

type BookStore struct {
	controller controllers.Controller
}

func NewService(controller controllers.Controller) *BookStore {
	return &BookStore{controller: controller}
}
