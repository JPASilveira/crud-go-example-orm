// Package services contains the business logic of the application.
package services

import (
	"crudOrm/internal/models"
	"crudOrm/internal/repositories"
)

// PersonService provides methods for person and address-related operations.
// It encapsulates the business logic and interacts with the repositories.
type PersonService struct {
	PersonRepo  repositories.PersonRepository
	AddressRepo repositories.AddressRepository
}

// CreatePersonWithAddresses creates a new person along with their associated addresses.
func (s PersonService) CreatePersonWithAddresses(person *models.Person) error {
	return s.PersonRepo.Create(person)
}

// GetPerson retrieves a person by their ID.
func (s PersonService) GetPerson(id uint) (*models.Person, error) {
	return s.PersonRepo.GetByID(id)
}

// GetAllPeople retrieves all people from the database.
func (s PersonService) GetAllPeople() ([]models.Person, error) {
	return s.PersonRepo.GetAll()
}

// UpdatePerson updates an existing person's information.
func (s PersonService) UpdatePerson(person *models.Person) error {
	return s.PersonRepo.Update(person)
}

// DeletePerson deletes a person by their ID.
func (s PersonService) DeletePerson(id uint) error {
	return s.PersonRepo.Delete(id)
}

// AddAddressToPerson adds a new address to an existing person.
func (s PersonService) AddAddressToPerson(personID uint, address *models.Address) error {
	address.PersonID = personID
	return s.AddressRepo.Create(address)
}

// GetAddressByID retrieves an address by its ID.
func (s PersonService) GetAddressByID(id uint) (*models.Address, error) {
	return s.AddressRepo.GetByID(id)
}

// UpdateAddress updates an existing address.
func (s PersonService) UpdateAddress(address *models.Address) error {
	return s.AddressRepo.Update(address)
}

// DeleteAddress deletes an address by its ID.
func (s PersonService) DeleteAddress(id uint) error {
	return s.AddressRepo.Delete(id)
}

// GetAddressesByPersonID retrieves all addresses associated with a specific person.
func (s PersonService) GetAddressesByPersonID(personID uint) ([]models.Address, error) {
	return s.AddressRepo.GetAllByPersonID(personID)
}
