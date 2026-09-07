// Menandai file sebagai titik masuk utama aplikasi.
package main

// Mengambil modul standar Go untuk urusan formatting dan cetak teks.
import (
	// Modul untuk mencetak pesan ke terminal
	"fmt"

	// Modul standar Go untuk penanganan log dan HTTP server
	"log"
	"net/http"

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
	// Menghubungkan aplikasi ke database.
	err := database.Connect()
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Membuat GraphQL server menggunakan schema dan resolver.
	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{},
			},
		),
	)

	// Endpoint GraphQL.
	http.Handle("/query", srv)

	// Menampilkan pesan ke terminal bahwa server GraphQL siap digunakan
	fmt.Println("GraphQL Server is running!")
	fmt.Println("GraphQL endpoint: http://localhost:8080/query")

	// Menjalankan HTTP server pada port 8080.
	log.Fatal(http.ListenAndServe(":8080", nil))
}