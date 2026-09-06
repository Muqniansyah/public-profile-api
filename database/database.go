package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"public-profile-api/config"
)

// Variabel global untuk menyimpan koneksi database agar bisa dipakai di file lain
var DB *sql.DB

// Fungsi untuk membuka dan memverifikasi koneksi ke database
func Connect() error {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)

	var err error

	// Fungsi untuk membuka dan memverifikasi koneksi ke database
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	// Menguji apakah koneksi ke database benar-benar aktif/terhubung
	err = DB.Ping()
	if err != nil {
		return err
	}

	fmt.Println("Database connected successfully!")

	// Mengembalikan nil (tanpa error) jika koneksi berhasil
	return nil
}