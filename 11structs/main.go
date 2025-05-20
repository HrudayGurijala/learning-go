package main

import "fmt"

func main() {
	fmt.Println("structs")
	//no inheritance , no super/ parent
	hruday := User{"hruday", "hruday@go.dev", true, 21}
	fmt.Println(hruday)
	fmt.Printf("details of hruday are : %+v \n", hruday)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
