package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://www.youtube.com/watch?v=cl7_ouTMFh0&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=26"

func main() {
	fmt.Println("URL handling")

	res, err := url.Parse(myurl)

	if err != nil {
		panic(err)
	}

	// fmt.Println(res.Scheme)
	// fmt.Println(res.Host)
	// fmt.Println(res.Path)
	// fmt.Println(res.Port())
	fmt.Println(res.RawQuery)

	qparams := res.Query()

	fmt.Println("quesry params are : ",qparams)
	fmt.Println(qparams["v"])

	for _,value := range qparams {
		fmt.Println("params are :", value)
	}
}
