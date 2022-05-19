package main

import "fmt"

func main() {
	fmt.Println("Welcome to array")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[2] = "Banana"
	fruitList[3] = "Cheeku"

	fmt.Println("List= ", fruitList)
	fmt.Println("Length= ", len(fruitList))

	var vegList = [3]string{"Potato", "Beans", "Ladyfinger"}
	fmt.Println("vegetables =", vegList)
	fmt.Println("Length= ", len(vegList))

}
