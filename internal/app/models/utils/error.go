package utils

import (
	"net/http"
)

type ErrorResponse struct {
	ErrorNum  int
	ErrorType string
}

func MessageError(w http.ResponseWriter, r *http.Request, code int, msg string) {
	msg_err := ErrorResponse{ErrorNum: code, ErrorType: msg}
	page := []string{"internal/app/views/templates/forum.html"}
	w.WriteHeader(code)
	ExecuteTemplate(w, page, msg_err)
}
