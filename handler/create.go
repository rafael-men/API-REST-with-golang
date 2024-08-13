package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rafael-men/rest-api-with-golang/model"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var todo model.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)

	if err != nil {
		log.Printf("Decode error :(")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	id, err := model.Insert(todo)

	var resp map[string]any

	if err != nil {
		resp = map[string]any {
			"Error":true,
			"Message": fmt.Sprintf("Insert Error",err),
		}
	}
	else {
		resp = map[string]any {
			"Error":false,
			"Message":"Todo successfully inserted"
		}
	}
	w.Header().Add("Content-type","application/json")
}
