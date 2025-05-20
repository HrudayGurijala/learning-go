package main

import "fmt"

func main() {
	fmt.Println("functions in go")
	greeter()

	result,message := proAdder(3,5,7,9,11)
	fmt.Println(result)
	fmt.Println(message)


}


func proAdder(values ...int)(int,string){
	total :=0

	for i := range values {
		total += values[i]
	}
	return total,"tun tun tun tun sahur"
}

func adder(val1 int, val2 int) int{
	return val1 + val2
}

func greeter(){
	fmt.Println("na mogga gug")
}