package util

import (
	"fmt"

	scribble "github.com/nanobox-io/golang-scribble"
	"github.com/spf13/viper"
)

// GetDatabaseDriver Method to get the scribble.Driver object
func GetDatabaseDriver() *scribble.Driver {
	// a new scribble driver, providing the directory where it will be writing to,
	// and a qualified logger if desired
	db, err := scribble.New(viper.GetString("database.path"), nil)
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	return db
}
