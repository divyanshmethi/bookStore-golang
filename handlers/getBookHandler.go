package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Todo: Set correct response codes for all cases
// Todo: Interconversion between layers
func (service *BookStore) GetBook(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-Type", "application/json")
	headerVars := mux.Vars(r)
	if headerVars == nil || headerVars["bookId"] == "" {
		w.WriteHeader(400)
		_ = json.NewEncoder(w).Encode("Invalid request")
		return
	}
	res, err := service.controller.GetBook(headerVars["bookId"])
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
