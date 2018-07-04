package main

import "fmt"

func main() {
	var smaller int
	var larger int

	fmt.Println("Enter a number")
	fmt.Scan(&smaller)

	fmt.Println("Enter a larger number")
	fmt.Scan(&larger)

	rem := larger % smaller

	fmt.Println(rem)
}
