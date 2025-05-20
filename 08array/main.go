package main

import "fmt"

func main(){
	fmt.Println("array")
	

	var carList [5]string
	var carList1 = [5]string{"ferrari", "bmw", "merc"}

	carList[0] = "ferrari"
	carList[2] = "bmw"
	carList[3] = "mclaren"

	fmt.Println("car list is :", carList)
	fmt.Println("car list length :", len(carList))
	fmt.Println("car list is", carList1)
	fmt.Println("car list length :", len(carList1))
}