package models

// Person menyimpan data profil dari tabel 'people'
type Person struct {
	ID      int
	Name    string
	Age     int
	Gender  string
	Job     string
	Phone   string
	Email   string
	Address string
	City    string
	Country string
	Hobbies []Hobby
}