package main

import "fmt"

func main(){
	scopeEx()
}

func scopeEx(){
	var i string = "string"

	for i:=0; i<1; i++ {
		fmt.Println(i)
	}
	for i := 0; i < 1; i++ {
		i := true
		fmt.Println(i)
	}

	fmt.Println(i)
}

