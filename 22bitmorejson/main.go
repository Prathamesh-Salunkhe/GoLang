package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON")
	//EncodeJson()
	DecodeJson()
}

// func EncodeJson() {
// 	lcoCourse := []course{
// 		{"ReactJS", 233, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
// 		{"Angular", 250, "LearnCodeOnline.in", "bcd123", []string{"google", "js"}},
// 		{"ReactJS", 233, "LearnCodeOnline.in", "xyz123", []string{"full-stack-dev", "js"}},
// 	}

// 	finalJson, err := json.MarshalIndent(lcoCourse, "", "\t")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("%s\n", finalJson)
// }

func DecodeJson() {
	jsonDataFormWeb := []byte(`
	{
		"coursename": "ReactJS",
        "Price": 233,
		"website": "LearnCodeOnline.in",
		
		"tags": ["full-stack-dev","js"],
	}
	`)

	var lcoCourse course
	checkValid := json.Valid(jsonDataFormWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFormWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}
}
