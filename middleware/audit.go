package middleware

import (
	"fmt"
	"net/http"
)

func Auditor() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				fmt.Println("enter audit middleware")
				next.ServeHTTP(w, r)
				fmt.Println("exit audit middleware")
			},
		)
	}
}
