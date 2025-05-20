package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("if-else")

	fmt.Print("Enter your age: ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	UserAge, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	var result string

	if UserAge < 10 {
		result = "invalid user"
	} else if UserAge > 60 {
		result = "invalid user gramps"
	} else {
		result = "valid user"
	}

	fmt.Println(result)

	if num := UserAge; num%2 == 0 {
		fmt.Println("num is even")
	}else {
		fmt.Println("num is odd")
	}

}
