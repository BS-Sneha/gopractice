package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the ratings from 1 to 5")
	input, _ := reader.ReadString('\n')
	//fmt.Printf("Type of input %T\n", input)
	fmt.Println("Thanks for the ratings", input)

}
