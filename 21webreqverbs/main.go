package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	//	"strings"
)

func main() {
	fmt.Println("Welcome to web verb")
	//	PerformGetRequest()
	//  PerformPostRequest()
	PerformPostFormRequest()
}

// func PerformGetRequest() {
// 	const myurl = "http://localhost:3000/get"

// 	response, err := http.Get(myurl)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer response.Body.Close()
// 	fmt.Println("Status code: ", response.StatusCode)
// 	fmt.Println("Content length is: ", response.ContentLength)

// 	var responseString strings.Builder
// 	content, _ := ioutil.ReadAll(response.Body)
// 	byteCount, _ := responseString.Write(content)
// 	fmt.Println("ByteCount is: ", byteCount)
// 	fmt.Println(responseString.String())
// }

// func PerformPostRequest() {
// 	const myurl = "http://localhost:3000/post"

// 	//fake json payload

// 	requestBody := strings.NewReader(`
//         {

//            "coursename":"GoLang",
// 	       "price":0,
// 	       "platform": "learnCodeOnline.in"
// 	    }
// 	`)
// 	response, err := http.Post(myurl, "application/json", requestBody)

// 	if err != nil {
// 		panic(err)
// 	}
// 	defer response.Body.Close()
// 	content, _ := ioutil.ReadAll(response.Body)
// 	fmt.Println(string(content))
// }

func PerformPostFormRequest() {
	const myurl = "http://localhost:3000/postform"
	//formdata
	data := url.Values{}
	data.Add("firstname", "Vipul")
	data.Add("lastname", "salunkhe")
	data.Add("email", "vvs@email.com")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
