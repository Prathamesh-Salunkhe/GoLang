package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Welcome to MouriTech")
	fmt.Println("Please Rate our pizza")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for Rating,", input)

	sample, error = strconv.ParseFloat(input, 64)

	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println("Added 1 to ur rating: ", sample+1)
	}
}
