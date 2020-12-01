package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	scribble "github.com/nanobox-io/golang-scribble"
	"github.com/spf13/viper"

	"github.com/tysonpaul89/go-gorilla-mux-example/middleware"
	"github.com/tysonpaul89/go-gorilla-mux-example/models"
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

	// ======================== Common Middlewares =============================
	// Getting router object.
	r := mux.NewRouter()

	// Adding middlewares
	r.Use(middleware.LoggingMiddleware)
	r.Use(middleware.JsonHeaderMiddleware)

	// ======================== Routes Definitions =============================
	bookObj := models.Book{}
	// Definding route and route method.
	r.HandleFunc("/books", bookObj.GetBooks).Methods("GET")
	r.HandleFunc("/book/{id}", bookObj.GetBook).Methods("GET")
	r.HandleFunc("/book", bookObj.CreateBook).Methods("POST")

	// ======================== Server Configs =================================
	// Runs the server in the 8000 port.
	// Use localhost:8000 to access the URL.
	log.Fatal(http.ListenAndServe(viper.GetString("port"), r))
}
