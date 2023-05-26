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

	var pegawai Employee
	var kantor Departement
	db.First(&pegawai, 1)
	fmt.Println(pegawai)
	db.First(&kantor, 2)
	fmt.Println(kantor)
}
