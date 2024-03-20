package handlers

import (
	"fmt"
	"net/http"
)

func ServerHealthHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body := r.Body
		defer body.Close()
        w.WriteHeader(http.StatusOK)
        fmt.Fprintf(w, "Server is running")
	}
}
