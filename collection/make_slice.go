package main

import "fmt"

func main() {
	slice := make([]int, 5, 10)
	for i := range slice {
		slice[i] = i
	}
	fmt.Println(slice)

	items := [...]int{10, 20, 30, 40, 50}

	for _, v := range items {
		v *= 2
	}
	fmt.Println(items)
}
