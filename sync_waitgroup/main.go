package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("worker %d started \n", i)
	//some task is happenning
	fmt.Printf("worker %d end \n", i)
}
func main() {
	fmt.Println("now i am learning sync_waitgroup")
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) //Increment the WaitGroup counter
		go worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("Workers task complete")
}
