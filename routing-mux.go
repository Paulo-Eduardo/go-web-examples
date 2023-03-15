package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func MuxRouting() {
	router := mux.NewRouter()
	router.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"] // the book title slug
		page := vars["page"] // the page

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	http.ListenAndServe(":80", router)
}

//
// Restrict for methods
//
// r.HandleFunc("/books/{title}", CreateBook).Methods("POST")
// r.HandleFunc("/books/{title}", ReadBook).Methods("GET")
// r.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
// r.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")

//
// Restrict for hostnames or subdomains
//
// r.HandleFunc("/books/{title}", BookHandler).Host("www.mybookstore.com")

//
// Restrict for schemes http/https
//
// r.HandleFunc("/secure", SecureHandler).Schemes("https")
// r.HandleFunc("/insecure", InsecureHandler).Schemes("http")

//
// Create new route for prefix path
//
// bookrouter := r.PathPreofix("/books").Subrouter()
// bookrouter.HandleFunc("/", AllBooks)
// bookrouter.HandleFunc("/{title}", GetBook)