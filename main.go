package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// Book Model
type Book struct {
	ID     uuid.UUID `json:"id"` // <-- Struct tags. Its represent what to display when its encoded as json. {"id": "as-12"}
	Title  string    `json:"title"`
	Author *Author   `json:"author"`
	Price  float64   `json:"price"`
}

// Author model
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var books []Book

func main() {
	// Getting router object.
	r := mux.NewRouter()

	// Definding route and route method.
	r.HandleFunc("/books", getBooks).Methods("GET")

	// Adding some sample data
	books = append(books, Book{ID: uuid.New(), Title: "Song of Ice and Fire", Price: 20.5, Author: &Author{Firstname: "George", Lastname: "R. R. Martin"}})
	books = append(books, Book{ID: uuid.New(), Title: "The Lord of the Rings", Price: 23.8, Author: &Author{Firstname: "J.R.R.", Lastname: "Tolkien"}})

	// Runs the server in the 8000 port.
	// Use localhost:8000 to access the URL.
	log.Fatal(http.ListenAndServe(":8000", r))
}

// Gets all the books
func getBooks(w http.ResponseWriter, r *http.Request) {
	// Setting response header
	w.Header().Set("Content-Type", "application/json")
	// This will convert books slice into a json string and writes into the http.ResponseWriter
	json.NewEncoder(w).Encode(books)
}
