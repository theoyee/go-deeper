package main

import (
	"fmt"
	"maps"
)

func main() {
	// p := map[string]int{
	// 	"alice": 4,
	// 	"ade":   5,
	// }

	m := make(map[string]int)

	m["one"] = 7
	m["two"] = 14

	fmt.Println("map: ", m)

	v1 := m["one"]
	fmt.Println("v1:", v1)
	v2 := m["two"]
	fmt.Println("v2:", v2)

	fmt.Println("len:", len(m))

	delete(m, "one")
	fmt.Println("map:", m)

	clear(m)
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
