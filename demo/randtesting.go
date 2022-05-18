package main

import (
	"fmt"
	"math/big"

	//"math/rand"
	//"time"
	"crypto/rand"
)

func main() {
	//rand.Seed(time.Now().UnixNano())
	//fmt.Println(rand.Intn(5) + 1)

	//random numbers uisng crypto

	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)

}
