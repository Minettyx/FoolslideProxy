package errors

import (
	"io"
	"net/http"
)

func NotFound(w http.ResponseWriter) {
	w.Header().Set("content-type", "text/plain; charset=utf-8")
	w.WriteHeader(404)
	io.WriteString(w, "404 Not Found")
}

func ServerError(w http.ResponseWriter) {
	w.Header().Set("content-type", "text/plain; charset=utf-8")
	w.WriteHeader(500)
	io.WriteString(w, "500 Internal Server Error")
}

func BadRequest(w http.ResponseWriter) {
	w.Header().Set("content-type", "text/plain; charset=utf-8")
	w.WriteHeader(400)
	io.WriteString(w, "400 Bad Request")
}
