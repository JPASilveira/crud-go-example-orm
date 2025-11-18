// Package models contains the data structures for the application.
package models

import "time"

// Person represents the structure of a person in the database.
type Person struct {
	// ID is the primary key for the Person model.
	ID uint `gorm:"primary_key" json:"id"`
	// FirstName is the person's first name.
	FirstName string `gorm:"not null;size:50" json:"first_name"`
	// LastName is the person's last name.
	LastName string `gorm:"not null;size:50" json:"last_name"`
	// Document is the person's unique document number.
	Document string `gorm:"unique not null;size:14;" json:"document"`
	// Phone is the person's phone number.
	Phone string `gorm:"not null;size:14" json:"phone"`
	// Email is the person's email address.
	Email string `gorm:"not null;size:100" json:"email"`
	// BirthDate is the person's date of birth.
	BirthDate time.Time `gorm:"not null;type:datetime" json:"birth_date"`
	// RegisterDate is the date the person was registered in the system.
	RegisterDate time.Time `gorm:"not null;type:datetime" json:"register_date"`
	// Addresses is a list of addresses associated with the person.
	// The `gorm` tag defines the foreign key relationship and cascade behavior.
	Addresses []Address `gorm:"foreignkey:PersonID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"addresses"`
}
