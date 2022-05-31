package main

import "fmt"

func main() {
	greeter()
	fmt.Println("Functions in GOLANG")

	result := adder(3, 4)
	fmt.Println("Addition = ", result)

	res := proAdder(3, 45, 5)
	fmt.Println("Addition = ", res)
}

func greeter() {
	fmt.Println("Nameastey ")
}

func adder(a int, b int) int {
	return a + b
}

func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total

}
