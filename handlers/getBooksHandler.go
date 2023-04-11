package handlers

import (
	"encoding/json"
	"net/http"
)

func (service *BookStore) GetBooks(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-Type", "application/json")
	res, err := service.controller.GetAllBooks()
	if err != nil {
		w.WriteHeader(500)
		_ = json.NewEncoder(w).Encode(err)
		return
	}
	if res == nil {
		_ = json.NewEncoder(w).Encode("No records found")
		return
	}
	_ = json.NewEncoder(w).Encode(res)
}
