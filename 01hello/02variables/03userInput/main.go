package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating: ")

	//comma ok //error ok
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating : ", input)
	fmt.Printf("Type of this rating= %T", input)
}
