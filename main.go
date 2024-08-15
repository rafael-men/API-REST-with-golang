package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rafael-men/rest-api-with-golang/config"
	"github.com/rafael-men/rest-api-with-golang/handler"
)

func main() {
	err := config.Load()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Post("/", handler.Create)
	r.Put("/{id}", handler.Update)
	r.Delete("/{id}", handler.Delete)
	r.Get("/", handler.List)
	r.Get("/", handler.Get)
	http.ListenAndServe(fmt.Sprintf(":%s", config.GetServerPort()), r)
}
