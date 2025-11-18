// Package models contains the data structures for the application.
package models

// Address represents the structure of an address in the database.
type Address struct {
	// ID is the primary key for the Address model.
	ID uint `gorm:"primary_key"`
	// ZipCode is the address's zip code.
	ZipCode string `gorm:"size:9" json:"zip_code"`
	// Street is the street name.
	Street string `gorm:"size:50" json:"street"`
	// Number is the building or house number.
	Number string `gorm:"size:10" json:"number"`
	// Complement is additional address information, such as an apartment number.
	Complement string `gorm:"size:30" json:"complement"`
	// Neighborhood is the neighborhood name.
	Neighborhood string `gorm:"size:50" json:"neighborhood"`
	// City is the city name.
	City string `gorm:"size:50" json:"city"`
	// State is the state or province.
	State string `gorm:"size:50" json:"state"`
	// PersonID is the foreign key that links the address to a person.
	PersonID uint
}
