package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Welcome to slices")
	var fruitList = []string{"Banana", "Mango", "Pear"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "Apple")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScore := make([]int, 4)

	highScore[0] = 235
	highScore[1] = 124
	highScore[2] = 878
	highScore[3] = 775
	//highScore[4] = 555
	highScore = append(highScore, 147, 258, 369)
	fmt.Println(highScore)

	sort.Ints(highScore)
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore))

	//how to remove a value from slices based on index

	var course = []string{"reacjs", "java", "python", "ruby", "js"}
	fmt.Println(course)
	var index int = 2
	course = append(course[:index], course[index+1:]...)
	fmt.Println(course)
}
