// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// )

// func main() {
// 	err := ioutil.WriteFile("test.txt", []byte("test data"), 0600)
// 	if err != nil {
// 		fmt.Println("error while creating file")
// 	}

// 	file, err := ioutil.ReadFile("test.txt")
// 	if err != nil {
// 		fmt.Println(err, "error while reading file")
// 	}

// 	fmt.Println(string(file))

// 	filesFromDir, err := ioutil.ReadDir(".")
// 	if err != nil{
// 		fmt.Println(err, "error while checking directory")
// 	}

//		for _, file := range filesFromDir {
//			fmt.Println(file.IsDir(), file.Mode(), file.Size(), file.Name())
//		}
//	}
package main

func minimumFromFour(a, b, c, d int) int {
    min := a
    if b<min {
        min=b
    }
    if min<c {
        min=c
    }
    if min<d {
        min=d
    }
    return min
}
