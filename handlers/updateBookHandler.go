package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (service *BookStore) UpdateBook(context *gin.Context) {
	context.Request.Header.Set("Content-Type", "application/json")
	var book Book
	if err := context.ShouldBindJSON(&book); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Unable to read the provided book Information"})
		return
	}
	controllerReq := toControllerReq(&book)
	rowsAffected, err := service.controller.UpdateBook(controllerReq)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if rowsAffected == 0 {
		context.String(http.StatusOK, "No records found to be updated")
		return
	}
	//Todo: Check if we can return the updated object
	context.String(http.StatusOK, "Book Updated Successfully")
}
