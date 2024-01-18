package main

import "fmt"

func main() {
	var f = Adder()
	fmt.Print(f(1), " - ")
	fmt.Print(f(20), " - ")
	fmt.Print(f(300))
}

/*
	将函数作为返回值
*/
func Adder() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}
