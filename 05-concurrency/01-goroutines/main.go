package main

import (
	"fmt"
	"time"
)

func say(message string) {
	for i := 0; i < 3; i++ {
		fmt.Println(message)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go say("Hello from goroutine")

	say("Hello from main")
}
