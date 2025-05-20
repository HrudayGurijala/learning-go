package main

import (
	"fmt"
	"sort"
)

//important
func main(){
	fmt.Println("slices")

	var carList = []string{"ferrari","bmw","audi"}
	fmt.Printf("type of carList is %T\n",carList)

	//append
	carList = append(carList, "cadillac")
	fmt.Println(carList)

	//range is [ , )
	//slicing in slice
	// carList = append(carList[1:2])  
	carList = append(carList[:2])
	// carList = append(carList[1:])
	fmt.Println(carList)

	//allocate and update
	highScores := make([]int,4)

	highScores[0] = 1
	highScores[1] = 2
	highScores[2] = 3
	highScores[3] = 4
	
	highScores = append(highScores, 123,234345,456)
	fmt.Println(highScores)

	//sorting a slice
	sort.Ints(highScores)
	fmt.Println(highScores)

	//how to remove a value from slice based on index
	var courses = []string{"react","go","rust","java"}
	fmt.Println(courses)

	var index int = 2

	courses = append(courses[:index], courses[index+1 :]...)
	fmt.Println(courses)

}