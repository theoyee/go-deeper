package main

import "fmt"

const (
	Todo = iota
	InProgress
	Done
	Cancelled
)

func main() {
	fmt.Println(Todo)
	fmt.Println(InProgress)
	fmt.Println(Done)
	fmt.Println(Cancelled)
}
