// Package server serves up the endpoints cache via http
package server

import (
	"io"
	"net/http"
)

// Health serves our health route
func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, "ok")
}
