package main

import "fmt"

func main() {
	var number int = 42

	var decimal float64 = float64(number)

	fmt.Printf("%T \n", number)
	fmt.Printf("%T \n", decimal)
}
