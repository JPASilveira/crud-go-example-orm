package models

import "time"

type Person struct {
	ID           uint      `gorm:"primary_key" json:"id"`
	FirstName    string    `gorm:"not null;size:50" json:"first_name"`
	LastName     string    `gorm:"not null;size:50" json:"last_name"`
	Document     string    `gorm:"unique not null;size:14;" json:"document"`
	Phone        string    `gorm:"not null;size:14" json:"phone"`
	Email        string    `gorm:"not null;size:100" json:"email"`
	BirthDate    time.Time `gorm:"not null;type:datetime" json:"birth_date"`
	RegisterDate time.Time `gorm:"not null;type:datetime" json:"register_date"`
	Addresses    []Address `gorm:"foreignkey:PersonID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"addresses"`
}
