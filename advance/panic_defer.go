package main

import "fmt"

func main() {
	f()
	fmt.Println("return normal")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover in f()")
		}
	}()
	fmt.Println("calling g()")
	g(0)
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("defer in g")
	fmt.Println("call in g", i)
	g(i + 1)
}
