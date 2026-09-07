package repository

import (
	"database/sql"
	"fmt"

	"public-profile-api/database"
	"public-profile-api/models"
)

// GetAllPeople menjalankan query untuk mengambil semua baris data dari tabel 'people'
func GetAllPeople() ([]models.Person, error) {
	// Menjalankan query SELECT ke database
	rows, err := database.DB.Query(`
		SELECT
			id,
			name,
			age,
			gender,
			job,
			phone,
			email,
			address,
			city,
			country
		FROM people
	`)

	// Mengembalikan error jika query gagal dieksekusi
	if err != nil {
		return nil, err
	}

	// Memastikan koneksi rows ditutup setelah fungsi selesai mengeksekusi data
	defer rows.Close()

	var people []models.Person

	// Loop untuk membaca setiap baris data hasil query
	for rows.Next() {
		var person models.Person

		// Memindahkan nilai dari kolom database ke dalam variabel struct person
		err := rows.Scan(
			&person.ID,
			&person.Name,
			&person.Age,
			&person.Gender,
			&person.Job,
			&person.Phone,
			&person.Email,
			&person.Address,
			&person.City,
			&person.Country,
		)

		// Jika proses pembacaan baris gagal, kembalikan error
		if err != nil {
			return nil, err
		}

		// Menambahkan data person ke dalam slice people
		people = append(people, person)
	}

	// Memeriksa apakah ada error saat proses perulangan baris data
	return people, rows.Err()
}

// GetPersonByID mengambil satu baris data person spesifik berdasarkan ID
func GetPersonByID(id int) (models.Person, error) {
	var person models.Person

	// Menjalankan query SELECT dengan parameter ID dan memetakan hasilnya ke struct person
	err := database.DB.QueryRow(`
		SELECT
			id,
			name,
			age,
			gender,
			job,
			phone,
			email,
			address,
			city,
			country
		FROM people
		WHERE id = ?
	`, id).Scan(
		&person.ID,
		&person.Name,
		&person.Age,
		&person.Gender,
		&person.Job,
		&person.Phone,
		&person.Email,
		&person.Address,
		&person.City,
		&person.Country,
	)

	// Mengecek apakah data person dengan ID tersebut tidak ditemukan.
	if err == sql.ErrNoRows {
		return models.Person{}, fmt.Errorf("person with ID %d not found", id)
	}

	// Mengecek apakah terjadi error lain saat query database.
	if err != nil {
		return models.Person{}, err
	}

	return person, nil
}

// GetRandomPerson mengambil satu baris data person secara acak dari database
func GetRandomPerson() (models.Person, error) {
	var person models.Person

	// Menjalankan query SELECT dengan urutan acak (ORDER BY RAND()) dan membatasi hanya 1 data
	err := database.DB.QueryRow(`
		SELECT
			id,
			name,
			age,
			gender,
			job,
			phone,
			email,
			address,
			city,
			country
		FROM people
		ORDER BY RAND()
		LIMIT 1
	`).Scan(
		&person.ID,
		&person.Name,
		&person.Age,
		&person.Gender,
		&person.Job,
		&person.Phone,
		&person.Email,
		&person.Address,
		&person.City,
		&person.Country,
	)

	// Mengembalikan struct kosong dan error jika query gagal
	if err != nil {
		return models.Person{}, err
	}

	return person, nil
}

// GetHobbiesByPersonID mengambil daftar hobi milik orang tertentu berdasarkan personID
func GetHobbiesByPersonID(personID int) ([]models.Hobby, error) {
	// Menjalankan query INNER JOIN antara tabel hobbies dan person_hobbies berdasarkan person_id
	rows, err := database.DB.Query(`
		SELECT
			h.id,
			h.name
		FROM hobbies h
		INNER JOIN person_hobbies ph
			ON h.id = ph.hobby_id
		WHERE ph.person_id = ?
	`, personID)

	// Mengembalikan error jika query gagal dieksekusi
	if err != nil {
		return nil, err
	}

	// Memastikan koneksi rows ditutup setelah fungsi selesai mengeksekusi data
	defer rows.Close()

	var hobbies []models.Hobby

	// Loop untuk membaca setiap baris hobi hasil query
	for rows.Next() {
		var hobby models.Hobby

		// Memindahkan nilai kolom id dan name ke struct hobby
		err := rows.Scan(
			&hobby.ID,
			&hobby.Name,
		)

		// Jika proses pembacaan baris gagal, kembalikan error
		if err != nil {
			return nil, err
		}

		// Menambahkan hobi ke dalam slice hobbies
		hobbies = append(hobbies, hobby)
	}

	// Memeriksa apakah ada error saat proses perulangan baris data
	return hobbies, rows.Err()
}