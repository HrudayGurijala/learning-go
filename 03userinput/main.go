package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter your name:")

	reader := bufio.NewReader(os.Stdin)
	name,_ := reader.ReadString('\n')
	fmt.Println("Hello", name)
}
