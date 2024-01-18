package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Printf("This is the %d iteration\n", i)
	}
	var a int = 1

	for a > 0 {
		a--
	}
	// range 循环
	str := "this is a range for string"

	for idx, v := range str {
		fmt.Println("idx =", idx, "value = ", string(v))
	}
	// 无限循环
	for {
	}
	fmt.Println("aaa")
}
