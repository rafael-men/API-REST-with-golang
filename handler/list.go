package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/rafael-men/rest-api-with-golang/model"
)

func List(w http.ResponseWriter, r *http.Request) {
	// Obter o ID da URL
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Printf("Error parsing id: %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	// Chamar model.GetAll com o ID
	todos, err := model.GetAll(id)
	if err != nil {
		log.Printf("Error parsing registers", err)
	}
	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(todos)

}
