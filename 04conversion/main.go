package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome to pizza app")
	fmt.Println("please rate our pizza between 1 to 5")

	reader := bufio.NewReader(os.Stdin)	
	rating,_ := reader.ReadString('\n')
	fmt.Println("your rating is,",rating)

	numRating, err := strconv.Atoi(strings.TrimSpace(rating))

	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println("Added 1 to your rating :", numRating+1)
	}

}
