package main

import "fmt"

func main() {
	a := 1
	b := 2

	_, second := swap2(a, b)
	fmt.Println(second)
}
func swap2(a, b int) (int, int) {
	return b, a
}
