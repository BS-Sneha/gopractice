package main

import "fmt"

type author struct {
	firstName string
	lastName  string
}

func (a author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

type blogpost struct {
	title   string
	content string
	bio     string
	author
}

func (b blogpost) details() {
	fmt.Println("Title:", b.title)
	fmt.Println("Content:", b.content)
	fmt.Println("Bio:", b.bio)
	fmt.Println("Author:", b.fullName())

}

func main() {
	author1 := author{
		"Martin",
		"Luther",
	}

	blogpost1 := blogpost{
		"Go programming",
		"Go does not support inheritance",
		"Go is easy",
		author1,
	}

	blogpost1.details()
}
