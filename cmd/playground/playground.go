package main

import (
	"fmt"
	"github.com/aakash-rajur/http/register"
	"net/http"
)

func main() {
	r := register.NewRegister().
		Add("/health", http.HandlerFunc(http.NotFound)).
		Add("/api/v2/books", http.HandlerFunc(http.NotFound)).
		Add("/api/v2/books/{bookId}", http.HandlerFunc(http.NotFound)).
		Add("/api/v2/users", http.HandlerFunc(http.NotFound)).
		Add("/api/v2/users/{userId}", http.HandlerFunc(http.NotFound)).
		Add("/api/v2/users/{userId}/books", http.HandlerFunc(http.NotFound)).
		Add("/api/v2/rpc/{service}/{method}", http.HandlerFunc(http.NotFound))

	entry, params, err := r.Find("/api/v2/users/10/books")

	if err != nil {
		panic(err)
	}

	fmt.Printf("entry: %+v\nparams: %+v\n", entry, params)
}
