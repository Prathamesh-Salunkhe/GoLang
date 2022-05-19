package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Time study in GoLang")
	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 Monday"))

	fmt.Println(presentTime.Format("02-01-2006 Monday"))

	fmt.Println(presentTime.Format("02-01-2006 Monday 15:04:05"))

}
