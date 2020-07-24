package handler

import (
	"github.com/go-chi/chi"
	"net/http"
)

func lists(router chi.Router) {
	router.Get("/", getAllItems)
}

func getAllItems(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("hi world"))
}