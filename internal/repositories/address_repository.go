package repositories

import (
	"crudOrm/internal/database"
	"crudOrm/internal/models"
)

type AddressRepository struct{}

func (r AddressRepository) Create(address *models.Address) error {
	return database.DB.Create(address).Error
}

func (r AddressRepository) GetByID(id uint) (*models.Address, error) {
	var address models.Address

	result := database.DB.First(&address, id)

	return &address, result.Error
}

func (r AddressRepository) GetAll() ([]models.Address, error) {
	var addresses []models.Address

	result := database.DB.Find(&addresses)

	return addresses, result.Error
}

func (r AddressRepository) GetAllByPersonID(personID uint) ([]models.Address, error) {
	var addresses []models.Address
	result := database.DB.Where("person_id = ?", personID).Find(&addresses)
	return addresses, result.Error
}

func (r AddressRepository) Update(address *models.Address) error {

	return database.DB.Save(address).Error
}

func (r AddressRepository) Delete(id uint) error {
	return database.DB.Delete(&models.Address{}, id).Error
}
