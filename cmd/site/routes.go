package main

import (
	"net/http"

	"nightcrawler086.blog/site/middleware"
)

func addRoutes(mux *http.ServeMux) {
	audit := middleware.Auditor()
	mux.Handle("GET /", audit(handleRoot()))
}
