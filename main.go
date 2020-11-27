package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"

	"github.com/tysonpaul89/go-gorilla-mux-example/middleware"
	"github.com/tysonpaul89/go-gorilla-mux-example/models"
)

var books []models.Book

func main() {
	// Read and loads the config data
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	// Getting router object.
	r := mux.NewRouter()

	// Adding middlewares
	r.Use(middleware.LoggingMiddleware)
	r.Use(middleware.JsonHeaderMiddleware)

	// Definding route and route method.
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/book/{id}", getBook).Methods("GET")

	// Adding some sample data
	books = append(books, models.Book{ID: uuid.UUID.String(uuid.New()), Title: "Song of Ice and Fire", Price: 20.5, Author: &models.Author{Firstname: "George", Lastname: "R. R. Martin"}})
	books = append(books, models.Book{ID: uuid.UUID.String(uuid.New()), Title: "The Lord of the Rings", Price: 23.8, Author: &models.Author{Firstname: "J.R.R.", Lastname: "Tolkien"}})

	// Runs the server in the 8000 port.
	// Use localhost:8000 to access the URL.
	log.Fatal(http.ListenAndServe(viper.GetString("port"), r))
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
	json.NewEncoder(w).Encode(&models.Book{})
}
