package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "Welcome to movie app"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the movie ratings from 1 to 5")
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for the ratings", input)

	numRatings, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Adding 1 to the ratings", numRatings+1)
	}

}
