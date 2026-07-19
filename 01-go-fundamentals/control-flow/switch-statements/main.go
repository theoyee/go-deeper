package main

import "fmt"

func main() {
	day := "Monday"

	switch day {
	case "Monday":
		fmt.Println("Start of the week")
	case "Friday":
		fmt.Println("Almost the weekend")
	case "Sunday":
		fmt.Println("Rest day")
	default:
		fmt.Println("Regular day")
	}
}
