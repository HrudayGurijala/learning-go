package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("go channels ")

	myCh := make(chan int,2) // specify the number of channel or dont 
	wg := &sync.WaitGroup{}
	
	// myCh <- 5
	// fmt.Println(<-myCh)

	wg.Add(2)
	//receive only -declare based on the chan arrow
	go func(wg *sync.WaitGroup, ch <-chan int){
		val,isChannelOpen := <-myCh

		fmt.Println(isChannelOpen)
		fmt.Println(val)

		fmt.Println(<-myCh)
		// fmt.Println(<-myCh)
		wg.Done()
	}(wg,myCh)
	//send only
	go func(wg *sync.WaitGroup, ch chan<- int){
		// myCh <- 5
		myCh <- 0
		close(myCh)
		wg.Done()
	}(wg,myCh)
	// go func(wg *sync.WaitGroup, ch chan int){}(wg,myCh)

	wg.Wait()
}
