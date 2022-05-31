package main

import "fmt"

func main() {

	fmt.Println("If Else in GoLang")

	loginCount := 23
	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Extra user"
	} else {
		result = "User"
	}
	fmt.Println(result)
}
