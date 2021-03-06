package models

import (
	// Build in packages
	"encoding/json"
	"net/http"

	// External in packages
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"

	// Internal in packages
	"github.com/tysonpaul89/go-gorilla-mux-example/util"
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

// GetBooks Gets all books.
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
	books := []Book{}
	for _, b := range data {
		book := Book{}
		// Converts each item into byte object.
		if err := json.Unmarshal([]byte(b), &book); err != nil {
			panic(err)
		}
		books = append(books, book)
	}

	// This will convert books slice into a json string and writes into the http.ResponseWriter
	json.NewEncoder(w).Encode(books)
}

// GetBook Get a given book by id.
func (b *Book) GetBook(w http.ResponseWriter, r *http.Request) {
	// Gets the database object
	db := util.GetDatabaseDriver()

	// Gets the parameters passed from the URL
	params := mux.Vars(r)

	// Gets an item from the database
	book := Book{}
	err := db.Read(viper.GetString("database.name"), params["id"], &book)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	json.NewEncoder(w).Encode(book)
}

// CreateBook Inserts new a book data into database.
func (b *Book) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	type Response struct {
		ID string `json:"id"`
	}

	// Response struct
	response := Response{}

	// Gets the database object
	db := util.GetDatabaseDriver()

	// Gets the book data from the request body
	err := json.NewDecoder(r.Body).Decode(&book)

	if (Book{}) != book { // Checks if the structure of the book matches its type
		book.ID = uuid.UUID.String(uuid.New())
		if err := db.Write(viper.GetString("database.name"), book.ID, book); err != nil {
			panic(err)
		}
		response.ID = book.ID
	} else { // Empty data handling
		w.WriteHeader(http.StatusNotFound)
	}

	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(response)
}

// UpdateBook Updates book data by id
func (b *Book) UpdateBook(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		ID      string `json:"id"`
		Message string `json:"message"`
	}

	// Response struct
	response := Response{}

	// Gets the database object
	db := util.GetDatabaseDriver()

	// Gets the parameters passed from the URL
	params := mux.Vars(r)

	// Gets an item from the database
	book := Book{}
	err := db.Read(viper.GetString("database.name"), params["id"], &book)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		response.Message = "Book not found!"
	} else {
		// Gets the book data from the request body
		newBook := Book{}
		err := json.NewDecoder(r.Body).Decode(&newBook)

		if err != nil {
			panic(err)
		}

		if (Book{}) != newBook { // Checks if the structure of the book matches its type
			newBook.ID = book.ID
			// Replaces the old data with new one.
			if err := db.Write(viper.GetString("database.name"), book.ID, newBook); err != nil {
				panic(err)
			}
			response.ID = book.ID
		} else { // Empty data handling
			w.WriteHeader(http.StatusNotFound)
			response.Message = "Book data received is incomplete."
		}
	}

	json.NewEncoder(w).Encode(response)
}

// DeleteBook Delete a book by id
func (b *Book) DeleteBook(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Message string `json:"message"`
	}

	// Response struct
	response := Response{}

	// Gets the database object
	db := util.GetDatabaseDriver()

	// Gets the parameters passed from the URL
	params := mux.Vars(r)

	// Gets an item from the database
	book := Book{}
	err := db.Read(viper.GetString("database.name"), params["id"], &book)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		response.Message = "Book not found!"
	} else {
		// Deletes the book
		if err := db.Delete(viper.GetString("database.name"), params["id"]); err != nil {
			panic(err)
		}

		response.Message = "Book deleted successfully."
	}

	json.NewEncoder(w).Encode(response)
}
