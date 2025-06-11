package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var signals = []string{"test"}

var wg sync.WaitGroup
var mut sync.Mutex

func main() {
	// fmt.Println("go routine")
	StartNow := time.Now() // works like a time since this check point we can use multiple like these btw
	// go greeter("bello")
	// greeter("minyun")

	websiteList := []string{
		"https://hruday.is-a.dev",
		"https://go.dev",
		"https://google.com",
		"https://github.com",
		"https://leetcode.com",
	}

	for _, website := range websiteList {
		go getStatusCode(website)
		wg.Add(1)
	}
	wg.Wait()

	fmt.Println(signals)

	fmt.Println("time took to run the program: ", time.Since(StartNow))
}

func getStatusCode(endpoint string) {
	res, err := http.Get(endpoint)

	defer wg.Done()
	if err != nil {
		fmt.Println("error pinging the website")
	}else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d statuscode for %s\n", res.StatusCode, endpoint)
	}
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(2 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }
