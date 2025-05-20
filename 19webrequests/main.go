package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://google.com"

func main() {
	fmt.Println("Making web request...")

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close() // Always defer closing the body right after checking for errors

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Response body:")
	fmt.Println(string(body))
}
