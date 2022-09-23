// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Person struct {
// 	ID     int
// 	Number string
// 	Year   int
// 	Student []Students
// }
// type Students struct {
// 	lastName   string
// 	firstName  string
// 	middleName string
// 	birthday   string
// 	address    string
// 	phone      string
// 	rating     []int
// }

// type myStruct struct {
// 	ID int
// 	Number string
// 	Year int
// 	Students []Name
// }

// type Name struct{
// 	Lastname string
// 	FirstName string
// 	MiddleName string
// 	Birthday string
// 	Address string
// 	Phone string
// 	Rating []int
// }

// func main() {
// 	d := []byte(`{
// 		"ID":134,
// 		"Number":"ИЛМ-1274",
// 		"Year":2,
// 		"Students":[
// 				"LastName":"Вещий",
// 				"FirstName":"Лифон",
// 				"MiddleName":"Вениаминович",
// 				"Birthday":"4апреля1970года",
// 				"Address":"632432,г.Тобольск,ул.Киевская,дом6,квартира23",
// 				"Phone":"+7(948)709-47-24",
// 				"Rating":[1,2,3]
// 		]
// 	}`)

// 	var s myStruct
// 	if  err := json.Unmarshal(d, &s); err != nil {
// 		fmt.Println("error while unmarshalling", err)
// 		return
// 	}
// 	fmt.Printf("%v", s)
// }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )
// type Message struct {
// 	Name string
// 	Email string
// 	PhoneNumber int
// 	StudyPlaces []Studying
// }

// type Studying struct {
// 	Name string
// 	Adress string
// 	GraduatedYear int
// }

// func main(){
// 	m := Message{
// 		Name: "John",
// 		Email: "john28@gmail.com",
// 		PhoneNumber: 12332343454,
// 		StudyPlaces: []Studying{
// 			{
// 				Name: "FSK",
// 				Adress: "Ferghana",
// 				GraduatedYear: 2021,
// 			},
// 		},
// 	}

// 	marsh, err := json.Marshal(m)
// 	if err != nil {
// 		fmt.Println("error while marshalling", err)
// 		return
// 	}
// 	fmt.Printf("%s", marsh)
// }

// ##################################################

package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name     string
	Email    string
	Status   bool
	Language []byte
}

func main1() {
	newUser := user{
		Name:     "Alex",
		Email:    "email@gmail.com",
		Status:   true,
		Language: []byte("ru"),
	}
	data, err := json.Marshal(newUser)
	if err != nil {
		panic(err)
	}

	newUser.Language = []byte("en")
	err = json.Unmarshal(data, &newUser)

	if err != nil {
		panic(err)
	}
	fmt.Println(string(newUser.Language))
}
