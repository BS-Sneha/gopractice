package main

import "fmt"

func main() {
	languages := make(map[string]string)
	languages["js"] = "javascript"
	languages["rs"] = "react"
	languages["rb"] = "ruby"
	languages["py"] = "python"

	//fmt.Println(languages)

	// for key, value := range languages {
	// 	fmt.Printf("Key is %v and value is %v\n", key, value)
	// }

	//fmt.Println(languages["js"])
	newLang := "rb"
	value, ok := languages[newLang]
	if ok == true {
		fmt.Println("Langauage found", value, newLang)
	} else {
		fmt.Println("key not found")
	}

}
