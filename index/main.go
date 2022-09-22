package main

import (
	"database/sql"
	"fmt"
	"log"

	faker "github.com/bxcodec/faker/v4"
	_ "github.com/lib/pq"
)

type Employee struct {
	ID          int
	FirstName   string
	Lastname    string
	Salary      float64
	PhoneNumber string
	UserName    string
}

func main() {
	conString := `user=postgres password=1234 dbname=test_db sslmode=disable`
	db, err := sql.Open("postgres", conString)
	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}
	defer db.Close()

	for i := 1; i <= 10000; i++ {
		employee := Employee{}
		err := faker.FakeData(&employee)
		if err != nil {
			log.Fatalf("failed to fake data: %v", err)
		}

		employee.FirstName = faker.FirstName()
		employee.Lastname = faker.LastName()
		employee.PhoneNumber = faker.Phonenumber()
		employee.UserName = faker.UserName()

		insertQuery := `
		insert into employees(
			firstname,
			lastname,
			salary,
			phone_number,
			username
		) values ($1, $2, $3, $4, $5)`
		_, err = db.Exec(
			insertQuery,
			employee.FirstName,
			employee.Lastname,
			employee.Salary,
			employee.PhoneNumber,
			employee.UserName,
		)
		if err != nil {
			log.Fatalf("failed to insert new employee: %v", err)
		}

		if i%1000 == 0 {
			fmt.Println("count=", i)
		}
	}
}
