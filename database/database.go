package database

import (
	"crypto/tls"
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"

	"public-profile-api/config"
)

// Variabel global untuk menyimpan koneksi database agar bisa dipakai di file lain
var DB *sql.DB

// Fungsi untuk membuka dan memverifikasi koneksi ke database
func Connect() error {
    // Mengatur konfigurasi TLS untuk koneksi MariaDB Cloud.
    tlsConfig := &tls.Config{
		MinVersion:         tls.VersionTLS12,
		InsecureSkipVerify: true,
	}

    // Mendaftarkan konfigurasi TLS ke driver MySQL/MariaDB.
    err := mysql.RegisterTLSConfig("cloud", tlsConfig)

    if err != nil {
        return err
    }

    // Membuat Data Source Name (DSN) untuk koneksi database.
    dsn := fmt.Sprintf(
        "%s:%s@tcp(%s:%s)/%s?tls=cloud",
        config.GetEnv("DB_USER", "root"),
		config.GetEnv("DB_PASSWORD", ""),
		config.GetEnv("DB_HOST", "127.0.0.1"),
		config.GetEnv("DB_PORT", "3306"),
		config.GetEnv("DB_NAME", "public_profile_api"),
    )

    // Membuka koneksi database.
    DB, err = sql.Open("mysql", dsn)

    if err != nil {
        return err
    }

    // Menguji apakah koneksi benar-benar berhasil.
    err = DB.Ping()

    if err != nil {
        return err
    }

    fmt.Println("Database connected successfully!")

    return nil
}

// Close menutup koneksi database jika koneksi sedang aktif.
func Close() error {

	// Mengecek apakah koneksi database tersedia.
	if DB != nil {

		// Menutup koneksi database.
		return DB.Close()
	}

	// Tidak ada error jika koneksi belum pernah dibuat.
	return nil
}