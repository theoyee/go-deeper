package main

import "fmt"

func main() {
	age := 22
	hasID := true

	if age >= 18 {
		fmt.Println("Age requirement passed")

		if hasID {
			fmt.Println("Access granted")
		}
	}
}
