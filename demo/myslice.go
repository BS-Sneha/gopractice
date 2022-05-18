package main

import (
	"fmt"
	"sort"
)

func main() {
	var vegList = []string{"Tomato", "Onion", "Cucumber", "Green Peas"}
	fmt.Println(vegList)

	vegList = append(vegList, "Radish", "Carrot")
	fmt.Println(vegList)

	scores := make([]int, 4)
	scores[0] = 60
	scores[1] = 20
	scores[2] = 30
	scores[3] = 40

	fmt.Println(scores)
	fmt.Println(sort.IntsAreSorted(scores))
	//sort.Ints(scores)
	//fmt.Println(scores)

	var course = []string{"react", "golang", "angular", "ruby"}
	fmt.Println(course)
	var index int = 2
	course = append(course[:index], course[index+1:]...)
	fmt.Println(course)

}
