// Package repositories contains the data access layer for the application.
package repositories

import (
	"crudOrm/internal/database"
	"crudOrm/internal/models"
)

// PersonRepository is a repository for the Person model.
// It provides methods for interacting with the person data in the database.
type PersonRepository struct{}

// Create creates a new person in the database.
func (r PersonRepository) Create(person *models.Person) error {
	return database.DB.Create(person).Error
}

// GetByID retrieves a person by their ID, preloading their associated addresses.
func (r PersonRepository) GetByID(id uint) (*models.Person, error) {
	var person models.Person
	// Preload("Addresses") tells GORM to also load the addresses associated with the person.
	result := database.DB.Preload("Addresses").First(&person, id)
	return &person, result.Error
}

// GetAll retrieves all people from the database, preloading their associated addresses.
func (r PersonRepository) GetAll() ([]models.Person, error) {
	var persons []models.Person
	// Preload("Addresses") tells GORM to also load the addresses associated with each person.
	result := database.DB.Preload("Addresses").Find(&persons)
	return persons, result.Error
}

// Update updates an existing person in the database.
func (r PersonRepository) Update(person *models.Person) error {
	return database.DB.Save(person).Error
}

// Delete removes a person from the database by their ID.
func (r PersonRepository) Delete(id uint) error {
	return database.DB.Delete(&models.Person{}, id).Error
}
