package main

import (
	"fmt"
	"time"
)

func worker(ch chan string) {
	time.Sleep(2 * time.Second)

	ch <- "Worker finished!"
}

func main() {
	ch := make(chan string)

	go worker(ch)

	fmt.Println("Waiting for worker...")

	message := <-ch

	fmt.Println(message)

	fmt.Println("Program finished.")
}
