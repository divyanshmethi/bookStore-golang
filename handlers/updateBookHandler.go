package handlers

import (
	"encoding/json"
	"net/http"
)

func (service *BookStore) UpdateBook(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-Type", "application/json")
	var book *Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		w.WriteHeader(400)
		_ = json.NewEncoder(w).Encode("Unable to read the provided book Information")
		return
	}
	defer r.Body.Close()
	controllerReq := toControllerReq(book)
	rowsAffected, err := service.controller.UpdateBook(controllerReq)
	if err != nil {
		w.WriteHeader(500)
		_ = json.NewEncoder(w).Encode(err)
		return
	}
	if rowsAffected == 0 {
		_ = json.NewEncoder(w).Encode("No records found to be updated")
		return
	}
	_ = json.NewEncoder(w).Encode("Book Updated Successfully")
}
