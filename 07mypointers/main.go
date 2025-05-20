package main

import "fmt"

func main() {
	fmt.Println("pointers")
	var myNum int = 23
	var myPtr = &myNum

	//& --> reference
	//* --> dereference



	fmt.Println("my value is :", myPtr)
	fmt.Println("my value is :", *myPtr)

	*myPtr = *myPtr *2	
	fmt.Println("my value is :", myPtr)
	fmt.Println("my value is :", *myPtr)
}
