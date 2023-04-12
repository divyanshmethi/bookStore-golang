package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Todo: Use gorm
// Todo: Check if we want to propogate the context to lower functions as well
func (service *BookStore) DeleteBook(context *gin.Context) {
	context.Request.Header.Set("Content-Type", "application/json")
	if context.Param("bookId") == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	rowsAffected, err := service.controller.DeleteBook(context.Param("bookId"))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid request"})
		return
	}
	if rowsAffected == 0 {
		context.String(http.StatusOK, "No records found to be deleted")
		return
	}
	//Todo: Check if want to return the deleted record or id of deleted record
	context.String(http.StatusOK, "Book Deleted Successfully")
}
