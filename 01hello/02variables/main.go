package main

import "fmt"

const Login string = "Vip" //Here the first letter of constant variable is  written capital to declare it as a public.
func main() {
	var username string = "Prathamesh"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.84589465879465
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	//default values and some aliases
	var a int //int is similar to uint32 or uint64
	fmt.Println(a)
	fmt.Printf("Variable is of type: %T \n", a)

	//implicit type
	var website = "Prathamesh"
	fmt.Println(website)

	//no var style
	numberOfUser := 300000.25
	fmt.Println(numberOfUser)
	fmt.Printf("Variable is of type: %T \n", numberOfUser)

	//printing constant
	fmt.Println(Login)
	fmt.Printf("Constant variable= %T\n", Login)
}
