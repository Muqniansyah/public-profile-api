// Menandai file sebagai titik masuk utama aplikasi.
package main

// Mengambil modul standar Go untuk urusan formatting dan cetak teks.
import (
	// Modul untuk mencetak pesan ke terminal
	"fmt"
	// Modul standar Go untuk berinteraksi dengan sistem operasi (misal: membaca Environment Variables)
	"os"

	// Modul standar Go untuk penanganan log dan HTTP server
	"log"
	"net/http"

	// Library pihak ketiga untuk membaca dan memuat file konfigurasi .env ke dalam environment
	"github.com/joho/godotenv"

	// Mengimpor package database buatan sendiri
	"public-profile-api/database"
	// Mengimpor resolver dan schema GraphQL buatan sendiri
	"public-profile-api/graph"
	"public-profile-api/graph/generated"

	// Mengimpor handler GraphQL dari library gqlgen
	"github.com/99designs/gqlgen/graphql/handler"
)

// Fungsi main adalah titik awal (entry point) eksekusi program.
func main() {
	// Memuat Environment Variables dari file .env.
	err := godotenv.Load()

	if err != nil {
		log.Println("Warning: .env file not found, using system Environment Variables.")
	}

	// Menghubungkan aplikasi ke database.
	err = database.Connect()
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Menutup koneksi database ketika aplikasi berhenti menggunakan anonymous function.
	defer func() {
		if err := database.Close(); err != nil {
			log.Println("Failed to close database:", err)
		}
	}()

	// Membuat GraphQL server menggunakan schema dan resolver.
	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{},
			},
		),
	)

	// Endpoint GraphQL dengan konfigurasi CORS.
	http.Handle(
		"/query",
		corsMiddleware(srv),
	)

	// Menampilkan pesan ke terminal bahwa server GraphQL siap digunakan
	fmt.Println("GraphQL Server is running!")
	fmt.Println("GraphQL endpoint: http://localhost:8080/query")

	// Mengambil port dari environment variable. (Vercel menyediakan PORT secara otomatis saat deployment)
	port := os.Getenv("PORT")

	// Menggunakan port 8080 jika aplikasi dijalankan secara lokal.
	if port == "" {

		port = "8080"

	}

	// Menjalankan HTTP server.
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// corsMiddleware mengatur izin akses Cross-Origin Resource Sharing (CORS).
func corsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Mengambil Origin dari request frontend.
        origin := r.Header.Get("Origin")

        // Daftar frontend yang diizinkan mengakses backend.
        allowedOrigins := map[string]bool{
            "http://localhost:5173": true,
            "https://public-profile-api.vercel.app": true,
        }

        // Mengizinkan Origin jika terdapat dalam daftar allowedOrigins.
        if allowedOrigins[origin] {

            w.Header().Set(
                "Access-Control-Allow-Origin",
                origin,
            )

        }

        // Mengizinkan method HTTP yang digunakan frontend.
        w.Header().Set(
            "Access-Control-Allow-Methods",
            "POST, GET, OPTIONS",
        )

        // Mengizinkan header Content-Type dari request fetch().
        w.Header().Set(
            "Access-Control-Allow-Headers",
            "Content-Type",
        )

        // Browser mengirim preflight request OPTIONS sebelum POST tertentu.
        if r.Method == http.MethodOptions {

            w.WriteHeader(http.StatusNoContent)

            return

        }

        // Meneruskan request ke GraphQL server.
        next.ServeHTTP(w, r)

    })

}