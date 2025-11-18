// Package repositories contains the data access layer for the application.
package repositories

import (
	"crudOrm/internal/database"
	"crudOrm/internal/models"
)

// AddressRepository is a repository for the Address model.
// It provides methods for interacting with the address data in the database.
type AddressRepository struct{}

// Create creates a new address in the database.
func (r AddressRepository) Create(address *models.Address) error {
	return database.DB.Create(address).Error
}

// GetByID retrieves an address by its ID.
func (r AddressRepository) GetByID(id uint) (*models.Address, error) {
	var address models.Address
	result := database.DB.First(&address, id)
	return &address, result.Error
}

// GetAll retrieves all addresses from the database.
func (r AddressRepository) GetAll() ([]models.Address, error) {
	var addresses []models.Address
	result := database.DB.Find(&addresses)
	return addresses, result.Error
}

// GetAllByPersonID retrieves all addresses associated with a specific person ID.
func (r AddressRepository) GetAllByPersonID(personID uint) ([]models.Address, error) {
	var addresses []models.Address
	result := database.DB.Where("person_id = ?", personID).Find(&addresses)
	return addresses, result.Error
}

// Update updates an existing address in the database.
func (r AddressRepository) Update(address *models.Address) error {
	return database.DB.Save(address).Error
}

// Delete removes an address from the database by its ID.
func (r AddressRepository) Delete(id uint) error {
	return database.DB.Delete(&models.Address{}, id).Error
}
