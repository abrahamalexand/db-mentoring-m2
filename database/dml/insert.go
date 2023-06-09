package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "bekasi99"
	dbname   = "test_db_camp"
)

func ConnectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`UPDATE employees SET salary = 1000000 WHERE id = 1`)

	if err != nil {
		panic(err)
	}

	fmt.Println("Update data success")

	// delete data to table `employees`
	_, err = db.Exec(`DELETE FROM employees WHERE id = 5`)
	if err != nil {
		panic(err)
	}

	fmt.Println("Delete data success")
}
