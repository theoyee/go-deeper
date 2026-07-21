package main

import "fmt"

func main() {
	age := 17

	if age >= 18 {
		fmt.Println("You can vote")
	} else {
		fmt.Println("You cannot vote yet")
	}
}
