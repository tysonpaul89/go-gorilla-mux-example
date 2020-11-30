package main

import (
	// "bytes"
	// "encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// "github.com/google/uuid"
	"github.com/gorilla/mux"
	scribble "github.com/nanobox-io/golang-scribble"
	"github.com/spf13/viper"

	"github.com/tysonpaul89/go-gorilla-mux-example/middleware"
	"github.com/tysonpaul89/go-gorilla-mux-example/models"
	"github.com/tysonpaul89/go-gorilla-mux-example/util"
)

var db *scribble.Driver

func main() {
	// ======================== Getting configurations =========================
	// Read and loads the config data
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	// // To Adding some sample data
	// books := []models.Book{}
	// books = append(books, models.Book{ID: uuid.UUID.String(uuid.New()), Title: "Song of Ice and Fire", Price: 20.5, Author: &models.Author{Firstname: "George", Lastname: "R. R. Martin"}})
	// books = append(books, models.Book{ID: uuid.UUID.String(uuid.New()), Title: "The Lord of the Rings", Price: 23.8, Author: &models.Author{Firstname: "J.R.R.", Lastname: "Tolkien"}})
	// db = getDatabaseDriver()
	// for _, book := range books {
	// 	fmt.Println(book)
	// 	db.Write(
	// 		viper.GetString("database.name"),
	// 		book.ID,
	// 		book,
	// 	)
	// }
	// ======================== Common Middlewares =============================
	// Getting router object.
	r := mux.NewRouter()

	// Adding middlewares
	r.Use(middleware.LoggingMiddleware)
	r.Use(middleware.JsonHeaderMiddleware)

	// ======================== Routes =========================================
	// Definding route and route method.
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/book/{id}", getBook).Methods("GET")

	// ======================== Server Configs =================================
	// Runs the server in the 8000 port.
	// Use localhost:8000 to access the URL.
	log.Fatal(http.ListenAndServe(viper.GetString("port"), r))
}

// Gets all the books
func getBooks(w http.ResponseWriter, r *http.Request) {
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

// Get a given book by id
func getBook(w http.ResponseWriter, r *http.Request) {
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
