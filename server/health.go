package server

import (
	"io"
	"net/http"
)

// Serve our health route
func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, "ok")
}
