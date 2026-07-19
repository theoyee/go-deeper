package main

import "fmt"

func main() {
	text := "Go 🔥"

	fmt.Println("String:", text)
	fmt.Println("Length in bytes:", len(text))

	fmt.Println("\nBytes:")
	for i := 0; i < len(text); i++ {
		fmt.Printf("%d ", text[i])
	}

	fmt.Println("\n\nRunes:")
	for _, character := range text {
		fmt.Printf("%c ", character)
	}
}
