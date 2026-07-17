package main

import "fmt"

var name = "Global"

func main() {
	name := "Function"

	if true {
		name := "Block"
		fmt.Println(name)
	}

	fmt.Println(name)
}
