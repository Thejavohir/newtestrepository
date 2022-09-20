package main

import (
	"encoding/json"
	"fmt"
)

// type Person struct {
// 	FirstName string
// 	LastName  string
// 	Age       int
// }

// type Adress struct {
// }

// func main() {
// 	a := Person{
// 		FirstName: "Elon",
// 		LastName:  "Musk",
// 		Age:       37,
// 	}

// 	data, err := json.Marshal(a)
// 	if err != nil {
// 		fmt.Println(err, "while marshaling")
// 	}

// 	fmt.Println(string(data))
// }

type Person struct {
	ID     int
	Number string
	Year   int
	Students 
}
type Students struct {
	lastName   string
	firstName  string
	middleName string
	birthday   string
	address    string
	phone      string
	rating     []int
}

type Average struct {
}

func main() {
	c := []byte(
		{
			"ID":134,
			"Number":"ИЛМ-1274",
			"Year":2,
			"Students":[
				{
					"LastName":"Вещий",
					"FirstName":"Лифон",
					"MiddleName":"Вениаминович",
					"Birthday":"4апреля1970года",
					"Address":"632432,г.Тобольск,ул.Киевская,дом6,квартира23",
					"Phone":"+7(948)709-47-24",
					"Rating":[1,2,3]
				},
			]
		}
	)
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err, "while marshalling")
	}

	fmt.Println(string(data))
}
