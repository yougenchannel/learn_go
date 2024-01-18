package main

import (
	"fmt"
	"math/rand"
)

func main() {
	switch result := getInt(); {
	case result < 5:
		fmt.Println("a = 2/ 3")
	case result > 5:
		fmt.Println("a > 5")
	default:
		fmt.Println("nothing")

	}
}
func getInt() int {
	return rand.Intn(10) + 1
}
