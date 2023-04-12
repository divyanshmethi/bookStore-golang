package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (service *BookStore) CreateBook(context *gin.Context) {
	context.Request.Header.Set("Content-Type", "application/json")
	var book Book
	if err := context.ShouldBindJSON(&book); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Unable to read the provided book Information"})
		return
	}
	controllerReq := toControllerReq(&book)
	err := service.controller.CreateBook(controllerReq)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.String(http.StatusCreated, "Book Added Successfully")
}
