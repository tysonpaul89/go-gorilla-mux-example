package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/tysonpaul89/go-gorilla-mux-example/middleware"
)

// Book Model
type Book struct {
	ID     string  `json:"id"` // <-- Struct tags. Its represent what to display when its encoded as json. {"id": "as-12"}
	Title  string  `json:"title"`
	Author *Author `json:"author"`
	Price  float64 `json:"price"`
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

	// Adding middlewares
	r.Use(middleware.LoggingMiddleware)
	r.Use(middleware.JsonHeaderMiddleware)

	// Definding route and route method.
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/book/{id}", getBook).Methods("GET")

	// Adding some sample data
	books = append(books, Book{ID: uuid.UUID.String(uuid.New()), Title: "Song of Ice and Fire", Price: 20.5, Author: &Author{Firstname: "George", Lastname: "R. R. Martin"}})
	books = append(books, Book{ID: uuid.UUID.String(uuid.New()), Title: "The Lord of the Rings", Price: 23.8, Author: &Author{Firstname: "J.R.R.", Lastname: "Tolkien"}})

	// Runs the server in the 8000 port.
	// Use localhost:8000 to access the URL.
	log.Fatal(http.ListenAndServe(":8000", r))
}

// Gets all the books
func getBooks(w http.ResponseWriter, r *http.Request) {
	// This will convert books slice into a json string and writes into the http.ResponseWriter
	json.NewEncoder(w).Encode(books)
}

// Get a given book by id
func getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, book := range books {
		if book.ID == params["id"] {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}
