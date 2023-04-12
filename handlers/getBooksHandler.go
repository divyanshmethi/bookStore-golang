package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (service *BookStore) GetBooks(context *gin.Context) {
	context.Request.Header.Set("Content-Type", "application/json")
	res, err := service.controller.GetAllBooks()
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
