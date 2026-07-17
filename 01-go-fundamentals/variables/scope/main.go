package main

import "fmt"

var packageName = "Go Deeper"

func main() {
	var localName = "theoyee"

	fmt.Println(packageName)
	fmt.Println(localName)

	if true {
		var blockName = "Inside Block"

		fmt.Println(blockName)
	}
}
