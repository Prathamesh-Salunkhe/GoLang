package main

import "fmt"

func main() {
	fmt.Println("Structs in GoLang")

	//no inheritance in golang; No super or parent
	prathamesh := User{"Prathamesh", "pvs@gmail.com", true, 23}
	fmt.Println(prathamesh)
	fmt.Printf("Prathamesh Details are: %+v\n", prathamesh)
	fmt.Printf("Name is %v and email is %v.", prathamesh.Name, prathamesh.Email, prathamesh.Age, prathamesh.Status)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
