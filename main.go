// Menandai file sebagai titik masuk utama aplikasi.
package main

// Mengambil modul standar Go untuk urusan formatting dan cetak teks.
import (
	// Modul untuk mencetak pesan ke terminal
	"fmt"

	// Mengimpor package database buatan sendiri
	"public-profile-api/database"

	// Mengimpor package service untuk mengakses logika bisnis
	"public-profile-api/service"
)

// Fungsi main adalah titik awal (entry point) eksekusi program.
func main() {
	// 1. Memanggil fungsi Connect() untuk menyambungkan ke database
	err := database.Connect()
	// Mengecek apakah koneksi database mengalami error
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}

	// 2. Mengambil seluruh data orang melalui layer service
	people, err := service.GetAllPeople()
	// Mengambil seluruh data orang melalui layer service
	if err != nil {
		fmt.Println("Failed to get people:", err)
		return
	}

	// Menampilkan daftar orang satu per satu ke terminal
	for _, person := range people {
		fmt.Println(person)
	}

	// 3. Mengambil data orang secara spesifik berdasarkan ID (contoh: ID 2)
	person, err := service.GetPersonByID(2)
	// Mengecek apakah terjadi error saat mengambil data orang berdasarkan ID
	if err != nil {
		fmt.Println("Failed to get person:", err)
		return
	}
	// Menampilkan data orang spesifik ke terminal
	fmt.Println(person)

	// 4. Mengambil satu data orang secara acak dari database
	randomPerson, err := service.GetRandomPerson()
	// Mengecek apakah terjadi error saat mengambil data orang acak
	if err != nil {
		fmt.Println("Failed to get random person:", err)
		return
	}

	// Menampilkan label dan data orang acak ke terminal
	fmt.Println("Random Person:")
	fmt.Println(randomPerson)

	fmt.Println("public Profile Api is running!")
}