package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Book struct defines the structure of a book
type Book struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	ISBN        string `json:"isbn"`
	Description string `json:"description,omitempty"`
}

// ToJSON converts a Book type to JSON
func (b Book) ToJSON() []byte {
	ToJSON, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return ToJSON
}

// FromJSON converts JSON to Book type
func FromJSON(data []byte) Book {
	book := Book{}
	if err := json.Unmarshal(data, &book); err != nil {
		panic(err)
	}
	return book
}

// Books is a map of ISBN to Book
var Books = map[string]Book{
	"92369823": Book{Title: "The Hitchhiker's Guide to the Galaxy", Author: "Douglas Adams", ISBN: "92369823"},
	"00003212": Book{Title: "Esio Trot", Author: "Roald Dahl", ISBN: "00003212"},
}

// BooksHandleFunc is used a HandleFunc for Books API
func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
	switch method := r.Method; method {
	case http.MethodGet:
		books := AllBooks()
		writeJSON(w, books)
	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		book := FromJSON(body)
		isbn, created := CreateBook(book)
		if created {
			w.Header().Add("Location", "/api/books/"+isbn)
			w.WriteHeader(http.StatusCreated)
		} else {
			w.WriteHeader(http.StatusConflict)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported HTTP method"))
	}
}

// BookHandleFunc is used a HandleFunc for Book by ID API
func BookHandleFunc(w http.ResponseWriter, r *http.Request) {
	isbn := r.URL.Path[len("/api/books/"):]
	switch method := r.Method; method {
	case http.MethodGet:
		if book, found := GetBook(isbn); found {
			writeJSON(w, book)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case http.MethodPut:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		book := FromJSON(body)
		exists := UpdateBook(isbn, book)
		if exists {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case http.MethodDelete:
		if _, found := GetBook(isbn); found {
			DeleteBook(isbn)
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported HTTP method"))
	}
}

// AllBooks returns a slice of all books
func AllBooks() []Book {
	books := []Book{}
	for _, book := range Books {
		books = append(books, book)
	}
	return books
}

// CreateBook adds a new book to the map
func CreateBook(b Book) (string, bool) {
	_, exists := Books[b.ISBN]
	if exists {
		return "", false
	}
	Books[b.ISBN] = b
	return b.ISBN, true
}

// UpdateBook updates a book
func UpdateBook(isbn string, book Book) bool {
	_, exists := Books[isbn]
	if exists {
		Books[isbn] = book
	}
	return exists
}

// GetBook returns a book by isbn
func GetBook(isbn string) (Book, bool) {
	book, ok := Books[isbn]
	return book, ok
}

// DeleteBook removes a book from the map by ISBN key
func DeleteBook(isbn string) {
	delete(Books, isbn)
}

func writeJSON(w http.ResponseWriter, i interface{}) {
	b, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
}
