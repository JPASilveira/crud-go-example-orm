// Package database handles the database connection.
package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB is a global variable that holds the database connection instance.
var DB *gorm.DB

// Connect initializes the database connection using GORM and an in-memory SQLite database.
// It sets the global DB variable to the new database connection.
func Connect() error {
	// gorm.Open creates a new database connection.
	// sqlite.Open("orm.db") specifies the SQLite driver and the database file name.
	db, err := gorm.Open(sqlite.Open("orm.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = db
	return nil
}
