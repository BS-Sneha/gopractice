package main

import "fmt"

func main() {
	// num := 10

	// if num%2 == 0 {
	// 	fmt.Println("Number is even")
	// } else {
	// 	fmt.Println("Number is odd")
	// }

	season := 3
	switch season {
	case 1:
		fmt.Println("Summer")

	case 2:
		fmt.Println("Winter")

	case 3:
		fmt.Println("Spring")

		//fallthrough
	case 4:
		fmt.Println("Rainy")

	default:
		fmt.Println("No season matched")
	}

}
