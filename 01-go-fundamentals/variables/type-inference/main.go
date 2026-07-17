package main

import "fmt"

func main() {
	name := "Emmanuel"
	age := 22
	height := 1.75
	isDeveloper := true

	fmt.Printf("Name: %T\n", name)
	fmt.Printf("Age: %T\n", age)
	fmt.Printf("Height: %T\n", height)
	fmt.Printf("Developer: %T\n", isDeveloper)
}
