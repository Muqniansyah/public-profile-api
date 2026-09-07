package service

import (
	// Import package models untuk mengakses struct data (seperti Person, Hobby, dll.)
	"public-profile-api/models"
	// Import package repository untuk mengakses fungsi query ke database
	"public-profile-api/repository"
)

// GetAllPeople mengambil seluruh data orang beserta daftar hobi masing-masing
func GetAllPeople() ([]models.Person, error) {
	// Mengambil daftar orang dari repository
	people, err := repository.GetAllPeople()
	if err != nil {
		return nil, err
	}

	// Loop setiap data person untuk menambahkan data hobi
	for i := range people {
		person, err := AddHobbies(people[i])
		if err != nil {
			return nil, err
		}

		// Memperbarui data person di slice dengan data yang sudah dilengkapi hobi
		people[i] = person
	}

	return people, nil

	// Kode lama (hanya mengambil data orang tanpa hobi): 
	// return repository.GetAllPeople()
}

// GetPersonByID mengambil data satu orang spesifik berdasarkan ID beserta daftar hobinya
func GetPersonByID(id int) (models.Person, error) {
	// Mengambil data person berdasarkan ID dari repository
	person, err := repository.GetPersonByID(id)
	if err != nil {
		return models.Person{}, err
	}

	// Menambahkan daftar hobi ke data person tersebut
	return AddHobbies(person)

	// Kode lama (hanya mengambil data orang tanpa hobi):
	// return repository.GetPersonByID(id)
}

// GetRandomPerson mengambil data satu orang secara acak beserta daftar hobinya
func GetRandomPerson() (models.Person, error) {
	// Mengambil data person acak dari repository
	person, err := repository.GetRandomPerson()
	if err != nil {
		return models.Person{}, err
	}

	// Menambahkan daftar hobi ke data person acak tersebut
	return AddHobbies(person)

	// Kode lama (hanya mengambil data orang tanpa hobi):
	// return repository.GetRandomPerson()
}

// AddHobbies adalah fungsi helper untuk mengambil hobi dari repository dan memasukkannya ke struct Person
func AddHobbies(person models.Person) (models.Person, error) {
	// Mengambil daftar hobi dari repository berdasarkan person.ID
	hobbies, err := repository.GetHobbiesByPersonID(person.ID)
	if err != nil {
		return models.Person{}, err
	}

	// Mengisi field Hobbies pada struct person
	person.Hobbies = hobbies

	return person, nil
}