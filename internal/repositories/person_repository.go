package repositories

import (
	"crudOrm/internal/database"
	"crudOrm/internal/models"
)

type PersonRepository struct{}

func (r PersonRepository) Create(person *models.Person) error {
	return database.DB.Create(person).Error
}

func (r PersonRepository) GetByID(id uint) (*models.Person, error) {
	var person models.Person

	result := database.DB.Preload("Addresses").First(&person, id)

	return &person, result.Error
}

func (r PersonRepository) GetAll() ([]models.Person, error) {
	var persons []models.Person

	result := database.DB.Preload("Addresses").Find(&persons)

	return persons, result.Error
}

func (r PersonRepository) Update(person *models.Person) error {
	return database.DB.Save(person).Error
}

func (r PersonRepository) Delete(id uint) error {
	return database.DB.Delete(&models.Person{}, id).Error
}
