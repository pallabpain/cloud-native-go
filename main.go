package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/pallabpain/cloud-native-go/api"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "The Books API!")
	})
	// All Books API
	http.HandleFunc("/api/books", api.BooksHandleFunc)
	// Book by ID API
	http.HandleFunc("/api/books/", api.BookHandleFunc)

	if err := http.ListenAndServe(port(), nil); err != nil {
		panic(err)
	}
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8000"
	}
	return ":" + port
}
