package main

import "fmt"

func main() {
	// 	for i := 1; i < 10; i++ {
	// 		if i == 5 {
	// 			//break
	// 			//continue
	// 			goto trip
	// 		}
	// 		//fmt.Println(i)
	// 	}

	// trip:
	// 	fmt.Println("Welcome to go to statements")

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}
	//fmt.Println(days)

	for _, day := range days {
		//fmt.Printf("index is %v and value is %v\n", index, day)
		fmt.Printf(" value is %v\n", day)
	}

}
