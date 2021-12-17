package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

func TheLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println(fmt.Sprintf("%s %s", r.Method, r.URL.Path))
		ctx := context.WithValue(r.Context(), "data", "data:test")
		next.ServeHTTP(rw, r.WithContext(ctx))
	})
}

func main() {
	r := http.NewServeMux()

	r.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		getContextData := r.Context().Value("data")
		fmt.Println(getContextData)
		io.WriteString(rw, `{"message": "hello world.."}`)
	})

	m := TheLogger(r)

	http.ListenAndServe(":8000", m)
}
