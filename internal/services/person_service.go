package services

import (
	"crudOrm/internal/models"
	"crudOrm/internal/repositories"
)

type PersonService struct {
	PersonRepo  repositories.PersonRepository
	AddressRepo repositories.AddressRepository
}

func (s PersonService) CreatePersonWithAddresses(person *models.Person) error {
	return s.PersonRepo.Create(person)
}

func (s PersonService) GetPerson(id uint) (*models.Person, error) {
	return s.PersonRepo.GetByID(id)
}

func (s PersonService) GetAllPeople() ([]models.Person, error) {
	return s.PersonRepo.GetAll()
}

func (s PersonService) UpdatePerson(person *models.Person) error {
	return s.PersonRepo.Update(person)
}

func (s PersonService) DeletePerson(id uint) error {
	return s.PersonRepo.Delete(id)
}

func (s PersonService) AddAddressToPerson(personID uint, address *models.Address) error {
	address.PersonID = personID
	return s.AddressRepo.Create(address)
}

func (s PersonService) GetAddressByID(id uint) (*models.Address, error) {
	return s.AddressRepo.GetByID(id)
}

func (s PersonService) UpdateAddress(address *models.Address) error {
	return s.AddressRepo.Update(address)
}

func (s PersonService) DeleteAddress(id uint) error {
	return s.AddressRepo.Delete(id)
}

func (s PersonService) GetAddressesByPersonID(personID uint) ([]models.Address, error) {
	return s.AddressRepo.GetAllByPersonID(personID)
}
