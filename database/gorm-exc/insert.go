package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Employee struct {
	ID             uint
	Name           string
	Address        string
	Age            uint8
	Birthdate      string
	Level          string
	ID_Departement uint
}

type Departement struct {
	ID               uint
	Nama_Departement string
}

func main() {
	dsn := "host=localhost user=postgres password=bekasi99 dbname=test_db_camp port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	pegawai := [5]Employee{{ID: 12345, Name: "Alexander", Address: "Taman Galaxy", Age: 23, Birthdate: "1999-12-16", Level: "Supervisor", ID_Departement: 54132},
		{ID: 12367, Name: "Umar", Address: "Taman Gigi", Age: 20, Birthdate: "2002-02-14", Level: "Staff", ID_Departement: 54133},
		{ID: 12374, Name: "Haryo", Address: "Taman Lawang", Age: 19, Birthdate: "2003-12-15", Level: "Supervisor", ID_Departement: 54131},
		{ID: 12345, Name: "Anggiet", Address: "Tambun", Age: 24, Birthdate: "1998-11-01", Level: "Staff", ID_Departement: 54130},
		{ID: 12345, Name: "John", Address: "Taman Galaxy", Age: 21, Birthdate: "2002-01-16", Level: "Supervisor", ID_Departement: 54147}}

	kantor := [5]Departement{{ID: 12345}}

	result := db.Create(pegawai)

	fmt.Println("Rows: ", result.RowsAffected)
}
