package main

import "fmt"

const untypedNumber = 10
const typedNumber int = 10

func main() {
	var a int = untypedNumber
	var b int = typedNumber

	fmt.Println(a)
	fmt.Println(b)
}
