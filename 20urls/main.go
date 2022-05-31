//Handling URLs in GoLang
package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {
	fmt.Println("URLs in GoLang") // "URLs in GoLang"
	fmt.Println(myurl)            // https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb

	//parsing the url
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)   //https
	fmt.Println(result.Host)     //lco.dev:3000
	fmt.Println(result.Path)     ///learn
	fmt.Println(result.Port())   //3000
	fmt.Println(result.RawQuery) // coursename=reactjs&paymentid=ghbj456ghb

	qparams := result.Query()
	fmt.Println("The type of query are: %T\n", qparams)
	fmt.Println(qparams["coursename"]) //reactjs

	for _, val := range qparams {
		fmt.Println("Param is : ", val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}
	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}
