package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(name string, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 5; i++ {
		fmt.Printf("%s: %d\n", name, i)
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go worker("Worker A", &wg)
	go worker("Worker B", &wg)

	wg.Wait()

	fmt.Println("All workers finished.")
}
