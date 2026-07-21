package main

import "fmt"

type UserID int

type Score = int

func main() {
	var userID UserID = 101
	var score Score = 95

	fmt.Println(userID)
	fmt.Println(score)
}