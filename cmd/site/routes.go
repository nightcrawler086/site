package main

import "net/http"

func addRoutes(*http.ServeMux) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
}
