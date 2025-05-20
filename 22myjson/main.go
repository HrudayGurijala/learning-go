package main

import (
	"encoding/json"
	"fmt"
)

type courses struct {
	Name     string `json:"courseName"` //aliases
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              //prevents sending this data
	Tags     []string `json:"tags,omitempty"` //dont send anything when there is null values
}

func main() {
	fmt.Println("json in go")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	hrudayCourses := []courses{
		{"react", 299, "udemy.com", "abc123", []string{"webdev", "server"}},
		{"reactnative", 199, "coursera.com", "hruday123", []string{"mobile", "ios"}},
		{"fastapi", 259, "udemy.com", "pass123", nil},
	}

	//package this data

	finalJson, err := json.MarshalIndent(hrudayCourses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s \n", finalJson)
}

func DecodeJson() {
	exampleJsonData := []byte(` {
			"courseName": "react",
			"Price": 299,
			"website": "udemy.com",
			"tags": ["webdev","server"]
        }
	`)

	var hrudayCourse courses 

	checkJson := json.Valid(exampleJsonData)

	if checkJson {
		fmt.Println("valid json")
		json.Unmarshal(exampleJsonData, &hrudayCourse)
		fmt.Printf("%#v\n", hrudayCourse)
	}else {
		fmt.Println("json was not valid")
	}
}
