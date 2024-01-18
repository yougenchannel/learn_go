package main

import "fmt"

// 空接口， 任何类型都实现了空接口
type Any interface{}

func main() {

	var a Any

	a = 5
	switch a.(type) {
	case string:
		fmt.Println(a)
	case int:
		fmt.Println(a)
	}
}
