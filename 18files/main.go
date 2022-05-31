package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files ")
	content := "NEW FILE IS CREATED"
	file, err := os.Create("./golang.txt") //create a new file and write above content on it

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}

	fmt.Println("length is : ", length)
	defer file.Close()
	readFile("./golang.txt") //Below function is called here

}

//To read the data from the file
func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("Text data inside the file is ", string(databyte))

}
``