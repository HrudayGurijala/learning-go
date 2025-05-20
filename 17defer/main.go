package main

import "fmt"

func main() {
	fmt.Println("defer in go")
	// defer basically cuts and paste the line at the end of the program 
	// defer works in a stack manner LIFO
	defer fmt.Println("this prints atlast")
	defer fmt.Println("this will be penultimate element in the stack")
	defer fmt.Println("this will be at the stack top")
	fmt.Println("random bullshit")
	fmt.Println("random bullshit")
	fmt.Println("random bullshit")
	fmt.Println("random bullshit")
	Deez()
}

func Deez(){
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
