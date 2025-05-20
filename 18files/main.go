package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("files in go")

	content := "khaleja re release on 30th May 2025"

	file,err := os.Create("./myfile.txt")
	checkNilErr(err)

	length, err := io.WriteString(file,content)

	checkNilErr(err)
	fmt.Println("length of the string is : ", length)
	defer file.Close()
	readFile("./myfile.txt")
}

func readFile(filename string) {
	data,err := ioutil.ReadFile(filename)

	checkNilErr(err)

	fmt.Println("text in the file is : ", string(data))
}

func checkNilErr(err error){
	if err != nil {
		panic(err)
	}
}