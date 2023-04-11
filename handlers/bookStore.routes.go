package handlers

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func (service *BookStore) RegisterRoutesAndStart() {
	r := mux.NewRouter()
	r.HandleFunc("/books", service.GetBooks).Methods(http.MethodGet)
	r.HandleFunc("/books", service.CreateBook).Methods(http.MethodPost)
	r.HandleFunc("/books/{bookId}", service.UpdateBook).Methods(http.MethodPut)
	r.HandleFunc("/books/{bookId}", service.DeleteBook).Methods(http.MethodDelete)
	r.HandleFunc("/books/{bookId}", service.GetBook).Methods(http.MethodGet)
	server := http.Server{
		Addr:              ":8081",
		Handler:           r,
		ReadHeaderTimeout: 5 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
