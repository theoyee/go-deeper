package main

import "fmt"

func main() {
	var number complex64 = 2 + 3i
	var largeNumber complex128 = 10 + 5i

	fmt.Println(number)
	fmt.Println(largeNumber)

	fmt.Println("Real part: ", real(number))
	fmt.Println("Imaginary part: ", imag(largeNumber))
}
