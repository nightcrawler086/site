package main

import (
	"net/http"
)

func handleRoot() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Hello from the root handler!"))
			return
		},
	)
}
