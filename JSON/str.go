package main

import (
	"encoding/json"
	"fmt"
)

type myStruct struct {
	Name string `json:"name"`
	Age int `json:"omitempty"`
	Status bool `json:"-"`
}

func main() {
	m := myStruct{Name: "John Connor", Age: 45, Status: false}

	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s", data) 
}
