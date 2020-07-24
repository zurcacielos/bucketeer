package handler

import (
	"github.com/go-chi/render"
	"net/http"
)

type ErrorResponse struct {
	Err error `json:"-"`
	StatusCode int `json:"-"`
	StatusText string `json:"status_text"`
	Message string `json:"message"`
}

var (
	ErrMethodNotAllowed = &ErrorResponse{StatusCode: 405, Message: "Method not allowed"}
	ErrNotFound = &ErrorResponse{StatusCode: 404, Message: "Resource not found"}
)

func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.StatusCode)
	return nil
}
