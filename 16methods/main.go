package main

import "fmt"

func main() {
	fmt.Println("Structs in GoLang")

	//no inheritance in golang; No super or parent
	prathamesh := User{"Prathamesh", "pvs@gmail.com", true, 23}
	fmt.Println(prathamesh)
	fmt.Printf("Prathamesh Details are: %+v\n", prathamesh)
	fmt.Printf("Name is %v and email is %v\n", prathamesh.Name, prathamesh.Email)
	prathamesh.GetStatus()
	prathamesh.GetAge()
	prathamesh.NewMail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active : ", u.Status) //returns the status of user
}

func (u User) GetAge() {
	fmt.Println("Age = ", u.Age) //returns the age of the user
}

func (u User) NewMail() {
	u.Email = "test.mail"
	fmt.Println("Email = ", u.Email) //here when we trying to add new data through method it will not reflected on the original value but changes in the copy of that type.

}
