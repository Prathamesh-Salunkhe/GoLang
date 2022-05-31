package main

import "fmt"

func main() {
	fmt.Println("Loops in GoLang")

	days := []string{"Sunday", "Monday", "Tuesday", "Thursday", "Friday", "Saturday"}
	fmt.Println(days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for j := range days {
	// 	fmt.Println(days[j])
	// }

	for index, day := range days {
		fmt.Printf("index is %v and value is %v\n", index, day)
	}

	roughValue := 1
	for roughValue < 10 {
		fmt.Println("Value is : ", roughValue)
		roughValue++
	}
}
