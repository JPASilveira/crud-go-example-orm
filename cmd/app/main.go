package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"crudOrm/internal/database"
	"crudOrm/internal/models"
	"crudOrm/internal/repositories"
	"crudOrm/internal/services"

	"gorm.io/gorm"
)

var personService services.PersonService
var reader *bufio.Reader

func main() {
	log.SetFlags(0)
	reader = bufio.NewReader(os.Stdin)

	if err := database.Connect(); err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	if err := database.DB.AutoMigrate(&models.Person{}, &models.Address{}); err != nil {
		log.Fatalf("Failed to run AutoMigrate: %v", err)
	}

	personService = services.PersonService{
		PersonRepo:  repositories.PersonRepository{},
		AddressRepo: repositories.AddressRepository{},
	}

	for {
		displayMainMenu()
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			personMenu()
		case "2":
			addressMenu()
		case "0":
			fmt.Println("\nExiting application. Goodbye!")
			return
		default:
			fmt.Println("\nInvalid option. Please try again.")
		}
	}
}

func readInput(prompt string) string {
	fmt.Printf("%s: ", prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func readID(prompt string) (uint, bool) {
	idStr := readInput(prompt)
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		fmt.Println("Invalid ID format.")
		return 0, false
	}
	return uint(id), true
}

func displayMainMenu() {
	fmt.Println("\n==================================")
	fmt.Println("         MAIN MENU")
	fmt.Println("==================================")
	fmt.Println("1. Person Operations (CRUD)")
	fmt.Println("2. Address Operations (Related CRUD)")
	fmt.Println("0. Exit")
	fmt.Print("Choose an option: ")
}

func personMenu() {
	for {
		fmt.Println("\n--- PERSON CRUD MENU ---")
		fmt.Println("1. Create Person")
		fmt.Println("2. Read Person by ID")
		fmt.Println("3. Update Person")
		fmt.Println("4. Delete Person")
		fmt.Println("5. List All People")
		fmt.Println("0. Back to Main Menu")
		fmt.Print("Choose an option: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			createPersonHandler()
		case "2":
			readPersonHandler()
		case "3":
			updatePersonHandler()
		case "4":
			deletePersonHandler()
		case "5":
			listAllPeopleHandler()
		case "0":
			return
		default:
			fmt.Println("\nInvalid option. Please try again.")
		}
	}
}

func createPersonHandler() {
	fmt.Println("\n--- Create New Person ---")

	firstName := readInput("First Name")
	lastName := readInput("Last Name")
	document := readInput("Document")
	phone := readInput("Phone")
	email := readInput("Email")
	birthDateStr := readInput("Birth Date (YYYY-MM-DD)")

	birthDate, err := time.Parse("2006-01-02", birthDateStr)
	if err != nil {
		fmt.Printf("Error parsing date: %v\n", err)
		return
	}

	addresses := []models.Address{}
	fmt.Println("\n--- Register Addresses (Type 'done' for ZipCode to finish) ---")

	for {
		zipCode := readInput("ZipCode")
		if strings.ToLower(zipCode) == "done" {
			break
		}

		street := readInput("Street")
		number := readInput("Number")
		complement := readInput("Complement")
		neighborhood := readInput("Neighborhood")
		city := readInput("City")
		state := readInput("State")

		addresses = append(addresses, models.Address{
			ZipCode:      zipCode,
			Street:       street,
			Number:       number,
			Complement:   complement,
			Neighborhood: neighborhood,
			City:         city,
			State:        state,
		})
	}

	newPerson := models.Person{
		FirstName:    firstName,
		LastName:     lastName,
		Document:     document,
		Phone:        phone,
		Email:        email,
		BirthDate:    birthDate,
		RegisterDate: time.Now(),
		Addresses:    addresses,
	}

	if err := personService.CreatePersonWithAddresses(&newPerson); err != nil {
		fmt.Printf("Error creating person: %v\n", err)
		return
	}
	fmt.Printf("Person created successfully! ID: %d, Addresses registered: %d\n", newPerson.ID, len(newPerson.Addresses))
}

func readPersonHandler() {
	id, ok := readID("Enter Person ID")
	if !ok {
		return
	}

	person, err := personService.GetPerson(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Printf("Person with ID %d not found.\n", id)
		} else {
			fmt.Printf("Error fetching person: %v\n", err)
		}
		return
	}

	fmt.Printf("Person Found: %s %s (ID: %d)\n", person.FirstName, person.LastName, person.ID)
	fmt.Printf("Email: %s, Document: %s, Phone: %s\n", person.Email, person.Document, person.Phone)
	fmt.Printf("Birth Date: %s\n", person.BirthDate.Format("2006-01-02"))
	fmt.Printf("Addresses (%d):\n", len(person.Addresses))
	for _, addr := range person.Addresses {
		fmt.Printf("		ID %d: %s, %s (%s). %s - %s (ZipCode: %s)\n",
			addr.ID, addr.Street, addr.Number, addr.Complement, addr.City, addr.State, addr.ZipCode)
	}
}

func updatePersonHandler() {
	fmt.Println("\n--- Update Person ---")
	id, ok := readID("Enter Person ID to update")
	if !ok {
		return
	}

	person, err := personService.GetPerson(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Printf("Person with ID %d not found.\n", id)
		} else {
			fmt.Printf("Error fetching person: %v\n", err)
		}
		return
	}

	fmt.Println("Enter new values (leave blank to keep current value):")

	newFirstName := readInput(fmt.Sprintf("New First Name (Current: %s)", person.FirstName))
	if newFirstName != "" {
		person.FirstName = newFirstName
	}

	newLastName := readInput(fmt.Sprintf("New Last Name (Current: %s)", person.LastName))
	if newLastName != "" {
		person.LastName = newLastName
	}

	newDocument := readInput(fmt.Sprintf("New Document (Current: %s)", person.Document))
	if newDocument != "" {
		person.Document = newDocument
	}

	newPhone := readInput(fmt.Sprintf("New Phone (Current: %s)", person.Phone))
	if newPhone != "" {
		person.Phone = newPhone
	}

	newEmail := readInput(fmt.Sprintf("New Email (Current: %s)", person.Email))
	if newEmail != "" {
		person.Email = newEmail
	}

	currentBirthDate := person.BirthDate.Format("2006-01-02")
	newBirthDateStr := readInput(fmt.Sprintf("New Birth Date (YYYY-MM-DD) (Current: %s)", currentBirthDate))
	if newBirthDateStr != "" {
		newBirthDate, err := time.Parse("2006-01-02", newBirthDateStr)
		if err != nil {
			fmt.Printf("Error parsing new date. Keeping original date: %v\n", err)
		} else {
			person.BirthDate = newBirthDate
		}
	}

	if err := personService.UpdatePerson(person); err != nil {
		fmt.Printf("Error updating person: %v\n", err)
		return
	}
	fmt.Printf("Person ID %d updated successfully.\n", person.ID)
}

func deletePersonHandler() {
	id, ok := readID("Enter Person ID to delete")
	if !ok {
		return
	}

	if err := personService.DeletePerson(id); err != nil {
		fmt.Printf("Error deleting person: %v\n", err)
		return
	}

	fmt.Printf("Person ID %d deleted successfully (and addresses via CASCADE).\n", id)
}

func listAllPeopleHandler() {
	fmt.Println("\n--- List All People ---")
	people, err := personService.GetAllPeople()
	if err != nil {
		fmt.Printf("Error fetching all people: %v\n", err)
		return
	}

	if len(people) == 0 {
		fmt.Println("No people found in the database.")
		return
	}

	for _, p := range people {
		fmt.Printf("[ID: %d] %s %s | Email: %s | Addresses: %d\n",
			p.ID, p.FirstName, p.LastName, p.Email, len(p.Addresses))
	}
}

func addressMenu() {
	for {
		fmt.Println("\n--- ADDRESS CRUD MENU ---")
		fmt.Println("1. Add Address to a Person (C)")
		fmt.Println("2. Read Address by ID (R)")
		fmt.Println("3. Update Address by ID (U)")
		fmt.Println("4. Delete Address by ID (D)")
		fmt.Println("5. List Addresses by Person ID")
		fmt.Println("0. Back to Main Menu")
		fmt.Print("Choose an option: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			addAddressHandler()
		case "2":
			readAddressHandler()
		case "3":
			updateAddressHandler()
		case "4":
			deleteAddressHandler()
		case "5":
			listAddressesByPersonHandler()
		case "0":
			return
		default:
			fmt.Println("\nInvalid option. Please try again.")
		}
	}
}

func addAddressHandler() {
	fmt.Println("\n--- Add New Address ---")
	personID, ok := readID("Enter Person ID to link this address")
	if !ok {
		return
	}

	if _, err := personService.GetPerson(personID); err != nil {
		fmt.Printf("Person ID %d not found. Cannot link address.\n", personID)
		return
	}

	zipCode := readInput("ZipCode")
	street := readInput("Street")
	number := readInput("Number")
	complement := readInput("Complement")
	neighborhood := readInput("Neighborhood")
	city := readInput("City")
	state := readInput("State")

	newAddress := models.Address{
		ZipCode:      zipCode,
		Street:       street,
		Number:       number,
		Complement:   complement,
		Neighborhood: neighborhood,
		City:         city,
		State:        state,
	}

	if err := personService.AddAddressToPerson(personID, &newAddress); err != nil {
		fmt.Printf("Error adding address: %v\n", err)
		return
	}
	fmt.Printf("Address created successfully! ID: %d, linked to Person ID: %d\n", newAddress.ID, personID)
}

func readAddressHandler() {
	id, ok := readID("Enter Address ID")
	if !ok {
		return
	}

	address, err := personService.GetAddressByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Printf("Address with ID %d not found.\n", id)
		} else {
			fmt.Printf("Error fetching address: %v\n", err)
		}
		return
	}

	fmt.Printf("Address Found (ID: %d, PersonID: %d):\n", address.ID, address.PersonID)
	fmt.Printf("-> Street: %s, Number: %s, Complement: %s\n", address.Street, address.Number, address.Complement)
	fmt.Printf("-> Location: %s - %s/%s\n", address.Neighborhood, address.City, address.State)
	fmt.Printf("-> ZipCode: %s\n", address.ZipCode)
}

func updateAddressHandler() {
	id, ok := readID("Enter Address ID to update")
	if !ok {
		return
	}

	address, err := personService.GetAddressByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Printf("Address with ID %d not found.\n", id)
		} else {
			fmt.Printf("Error fetching address: %v\n", err)
		}
		return
	}

	fmt.Println("Enter new values (leave blank to keep current value):")

	zipCode := readInput(fmt.Sprintf("New ZipCode (Current: %s)", address.ZipCode))
	if zipCode != "" {
		address.ZipCode = zipCode
	}

	street := readInput(fmt.Sprintf("New Street (Current: %s)", address.Street))
	if street != "" {
		address.Street = street
	}

	number := readInput(fmt.Sprintf("New Number (Current: %s)", address.Number))
	if number != "" {
		address.Number = number
	}

	complement := readInput(fmt.Sprintf("New Complement (Current: %s)", address.Complement))
	if complement != "" {
		address.Complement = complement
	}

	neighborhood := readInput(fmt.Sprintf("New Neighborhood (Current: %s)", address.Neighborhood))
	if neighborhood != "" {
		address.Neighborhood = neighborhood
	}

	city := readInput(fmt.Sprintf("New City (Current: %s)", address.City))
	if city != "" {
		address.City = city
	}

	state := readInput(fmt.Sprintf("New State (Current: %s)", address.State))
	if state != "" {
		address.State = state
	}

	if err := personService.UpdateAddress(address); err != nil {
		fmt.Printf("Error updating address: %v\n", err)
		return
	}
	fmt.Printf("Address ID %d updated successfully.\n", address.ID)
}

func deleteAddressHandler() {
	id, ok := readID("Enter Address ID to delete")
	if !ok {
		return
	}

	if err := personService.DeleteAddress(id); err != nil {
		fmt.Printf("Error deleting address: %v\n", err)
		return
	}

	fmt.Printf("Address ID %d deleted successfully.\n", id)
}

func listAddressesByPersonHandler() {
	personID, ok := readID("Enter Person ID to list addresses")
	if !ok {
		return
	}

	addresses, err := personService.GetAddressesByPersonID(personID)
	if err != nil {
		fmt.Printf("Error fetching addresses: %v\n", err)
		return
	}

	if len(addresses) == 0 {
		fmt.Printf("Person ID %d has no addresses.\n", personID)
		return
	}

	fmt.Printf("Addresses for Person ID %d:\n", personID)
	for i, addr := range addresses {
		fmt.Printf("  [%d] ID %d: %s, %s (%s). %s - %s (ZipCode: %s)\n",
			i+1, addr.ID, addr.Street, addr.Number, addr.Complement, addr.City, addr.State, addr.ZipCode)
	}
}
