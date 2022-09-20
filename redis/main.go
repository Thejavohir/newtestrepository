package main

import (
	"fmt"
	"time"

	"github.com/"
)

type User struct {
	FirstName, LastName, Email, Password, Birthday string
	Age int
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		Password: "",
		DB: 0,
	})

	user := &User{
		FirstName: "firstName",
		LastName: "lastName",
		Email: "email",
		Birthday: "2002-20-02",
		Age: 20,
	}
	user2:=&User{
		FirstName: "firstName2",
		LastName: "lastName2",
		Email: "email2",
		Birthday: "2002-20-01",
		Age: 23,
	}




	err = client.Set("user", string(val), time.Minute*12).Err()

	if err != nil {
		fmt.Println(err, "error while setting json data to redis db")
	}

	value, err := client.Get("user").Result()
	if err != nil {
		fmt.Println(err)
	}

	user1:= User{}

}