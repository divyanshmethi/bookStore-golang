package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (service *BookStore) DeleteBook(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-Type", "application/json")
	headerVars := mux.Vars(r)
	if headerVars == nil || headerVars["bookId"] == "" {
		w.WriteHeader(400)
		_ = json.NewEncoder(w).Encode("Invalid request")
		return
	}
	rowsAffected, err := service.controller.DeleteBook(headerVars["bookId"])
	if err != nil {
		w.WriteHeader(500)
		_ = json.NewEncoder(w).Encode(err)
		return
	}
	if rowsAffected == 0 {
		_ = json.NewEncoder(w).Encode("No records found to be deleted")
		return
	}
	_ = json.NewEncoder(w).Encode("Book Deleted Successfully")
}
