package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (service *BookStore) RegisterRoutesAndStart() {
	server := gin.New()
	server.GET("/books", service.GetBooks)
	server.GET("/books", service.CreateBook)
	server.GET("/books/{bookId}", service.UpdateBook)
	server.GET("/books/{bookId}", service.DeleteBook)
	server.GET("/books/{bookId}", service.GetBook)
	if err := server.Run(":8081"); err != nil {
		panic(err)
	}
	//Todo: If we can pass custom context to gin
	fmt.Println("Server running on port 8081")
}
