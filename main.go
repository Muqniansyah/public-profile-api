// Menandai file sebagai titik masuk utama aplikasi.
package main

// Mengambil modul standar Go untuk urusan formatting dan cetak teks.
import (
	// Modul untuk mencetak pesan ke terminal
	"fmt"

	// Mengimpor package database buatan sendiri
	"public-profile-api/database"
)

// Fungsi main adalah titik awal (entry point) eksekusi program.
func main() {
	// Memanggil fungsi Connect() untuk menyambungkan ke database
	err := database.Connect()
	// Mengecek apakah koneksi database mengalami error
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}

	fmt.Println("public Profile Api is running!")
}