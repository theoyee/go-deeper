// package main

// import "fmt"

// func main() {
// 	var letter byte = 65

// 	fmt.Println(letter)
// 	fmt.Printf("%c\n", letter)
// }

package main

import "fmt"

func main() {
	// text := "🔥"
	text := "Go 🔥 Nigeria 🇳🇬"

	fmt.Println(len(text))

	for _, character := range text {
		fmt.Printf("%c\n", character)
	}
}
