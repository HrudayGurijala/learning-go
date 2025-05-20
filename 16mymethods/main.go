package main

import "fmt"

func main() {
	fmt.Println("methods")
	//no inheritance , no super/ parent
	hruday := User{"hruday", "hruday@go.dev", true, 21}
	fmt.Println(hruday)
	fmt.Printf("details of hruday are : %+v \n", hruday)
	hruday.GetStatus()
	hruday.NewMail()
	fmt.Printf("details of hruday are : %+v \n", hruday)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus(){
	fmt.Println("Is user active?", u.Status)

}

func (u User) NewMail(){
	u.Email = "googoo@go.dev"
	fmt.Println("Email of this user is :", u.Email)
}