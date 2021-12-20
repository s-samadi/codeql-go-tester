package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

// Validate the ownership of the ID
// Header "Authorization: ID" matches the supplied path ID
// e.g. curl -v localhost:8000/account/123 -H "Authorization: 123"
// In a real-world implementation, "Authorization: ID" would be a JWT claim
func AuthorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		profile := r.Header.Get("Authorization")
		if len(profile) == 0 {
			fmt.Println("missing auth token")
			rw.WriteHeader(401)
			return
		}
		tokenID := mux.Vars(r)["id"]
		if profile != tokenID {
			fmt.Println("ownership not matched")
			rw.WriteHeader(401)
			return
		}
		next.ServeHTTP(rw, r)
	})
}

func GetAccount(rw http.ResponseWriter, r *http.Request) {
	io.WriteString(rw, `{"message": "hello world.."}`)
}

func main() {
	fmt.Println("running...")
	router := mux.NewRouter()
	router.Handle("/account/{id}", AuthorizationMiddleware(http.HandlerFunc(GetAccount)))
	http.Handle("/", router)
	http.ListenAndServe(":8000", router)
}
