package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	// Use some middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Define your routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	r.Get("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is the About page.")
	})

	// Mount the router and start the server
	fmt.Println("Starting server on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
