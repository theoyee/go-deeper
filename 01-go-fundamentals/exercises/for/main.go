package main

import "fmt"

func main() {

	for i := range 3 {
		fmt.Println("range", i)
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
