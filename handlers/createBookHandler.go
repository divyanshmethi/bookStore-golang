package handlers

import (
	"encoding/json"
	"net/http"
)

func (service *BookStore) CreateBook(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-Type", "application/json")
	var book *Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		w.WriteHeader(400)
		_ = json.NewEncoder(w).Encode("Unable to read the provided book Information")
		return
	}
	defer r.Body.Close()
	controllerReq := toControllerReq(book)
	err := service.controller.CreateBook(controllerReq)
	if err != nil {
		w.WriteHeader(500)
		_ = json.NewEncoder(w).Encode(err)
		return
	}
	_ = json.NewEncoder(w).Encode("Book Added Successfully")
}
