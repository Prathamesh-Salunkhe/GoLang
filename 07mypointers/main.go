package main

import "fmt"

func main() {
	fmt.Println("Welcome to Pointers")

	// var ptr *int
	// fmt.Println("Value of Pointer= ", ptr)

	myNumber := 23

	var ptr = &myNumber
	fmt.Println("Value= ", ptr)
	fmt.Println("Value= ", *ptr)

	*ptr = *ptr + 2
	fmt.Println("Value= ", myNumber)
	fmt.Println("Value= ", ptr)
	fmt.Println("Value= ", *ptr)
	fmt.Println("Value= ", &myNumber)
}
