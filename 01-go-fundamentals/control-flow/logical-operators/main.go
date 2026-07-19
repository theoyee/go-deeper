package main

import "fmt"

func main() {
	age := 22
	hasID := true

	if age >= 18 && hasID {
		fmt.Println("Access granted")
	}
}
