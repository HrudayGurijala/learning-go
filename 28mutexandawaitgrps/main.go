package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("mutex and await groups")

	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	var score = []int{0}

	wg.Add(3) //specify the number of routines we are running
	go func(wg *sync.WaitGroup, mut *sync.Mutex){
		fmt.Println("first routine")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg,mut)

	go func(wg *sync.WaitGroup, mut *sync.Mutex){
		fmt.Println("second routine")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg,mut)

	go func(wg *sync.WaitGroup, mut *sync.Mutex){
		fmt.Println("third routine")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg,mut)

	wg.Wait()
	fmt.Println(score)
}
