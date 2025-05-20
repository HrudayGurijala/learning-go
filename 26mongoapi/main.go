package main

import (
	"fmt"
	"log"
	"mongoapi/router"
	"net/http"
)

func main() {
	fmt.Println("mongodb api")
	fmt.Println("server getting started ...")
	log.Fatal(http.ListenAndServe(":8000",router.Router()))
	fmt.Println("Listening to port 4000...")
}
