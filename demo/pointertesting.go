package main

import (
	"fmt"
)

func main() {
	//var ptr *int
	nos := 255
	var ptr = &nos

	fmt.Println("Address of ptr", ptr)
	fmt.Println("value of ptr", *ptr)

}
