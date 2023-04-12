package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (service *BookStore) GetBook(context *gin.Context) {
	context.Request.Header.Set("Content-Type", "application/json")
	if context.Param("bookId") == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	res, err := service.controller.GetBook(context.Param("bookId"))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if res == nil {
		context.String(http.StatusOK, "No records found")
		return
	}
	context.JSON(http.StatusOK, res)
}
