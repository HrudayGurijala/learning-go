package main

import "fmt"

func main() {
	fmt.Println("mymaps")

	languages := make(map[string]string)

	languages["js"] = "yavascript"
	languages["rb"] = "ruby"
	languages["py"] = "python"

	fmt.Println("list of languages", languages)
	fmt.Println("rb is short for :", languages["rb"])

	//deleting
	delete(languages,"rb")
	fmt.Println("list of languages",languages)

	//loops 
	for key, value := range languages {
	// for _, value := range languages { if you want only values
		fmt.Printf("for key %v , the value is %v\n", key, value)
	}
}
