package controller

import (
	// Build in packages
	"encoding/json"
	"net/http"

	// External in packages
	// "github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"

	// Internal in packages
	"github.com/tysonpaul89/go-gorilla-mux-example/models"
	"github.com/tysonpaul89/go-gorilla-mux-example/util"
)

// Book Object
type Book struct{}

// GetBooks Gets all books
func (b *Book) GetBooks(w http.ResponseWriter, r *http.Request) {
	// Gets the database object
	db := util.GetDatabaseDriver()

	// Get all the data from the database.
	data, err := db.ReadAll(
		viper.GetString("database.name"),
	)

	if err != nil {
		panic(err)
	}

	// Gets all items from database and appends it to books slice.
	books := []models.Book{}
	for _, b := range data {
		book := models.Book{}
		// Converts each item into byte object.
		if err := json.Unmarshal([]byte(b), &book); err != nil {
			panic(err)
		}
		books = append(books, book)
	}

	// This will convert books slice into a json string and writes into the http.ResponseWriter
	json.NewEncoder(w).Encode(books)
}

// GetBook Get a given book by id
func (b *Book) GetBook(w http.ResponseWriter, r *http.Request) {
	// Gets the database object
	db := util.GetDatabaseDriver()

	// Gets the parameters passed from the URL
	params := mux.Vars(r)

	// Gets an item from the database
	book := models.Book{}
	err := db.Read(viper.GetString("database.name"), params["id"], &book)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(book)
}
