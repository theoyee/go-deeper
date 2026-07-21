// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	var wg sync.WaitGroup

// 	wg.Add(3)

// 	go func() {
// 		defer wg.Done()
// 		fmt.Println("One")
// 	}()

// 	go func() {
// 		defer wg.Done()
// 		fmt.Println("Two")
// 	}()

// 	go func() {
// 		defer wg.Done()
// 		fmt.Println("Three")
// 	}()

//		wg.Wait()
//	}
