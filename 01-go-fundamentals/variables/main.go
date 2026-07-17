package main

import "fmt"

func main() {
	// Variable declaration with var
	var name string = "theoyee"

	// Type inference, infer types based on the valueof the variable
	var age = 23

	// Short variable declaration
	city := "Abuja"

	// Reassignment
	age = 24

	// Variables in Go are statically typed.
	// Once Go determines the type of a variable, that variable cannot hold a completely different type.

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(city)
}
