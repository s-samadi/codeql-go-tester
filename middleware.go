// Basic authorization middleware

package main

import (
	"fmt"
	"net/http"
)

// Hardcoded auth token
var TOKEN = "123"

// Header "Authorization: ID" matches the TOKEN
// In a real-world implementation, "Authorization: ID" would be a JWT claim
func AuthorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		id := r.Header.Get("Authorization")
		if len(id) == 0 {
			fmt.Println("missing auth token")
			rw.WriteHeader(http.StatusUnauthorized)
			return
		}
		if id != TOKEN {
			fmt.Println("ownership not matched")
			rw.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(rw, r)
	})
}
