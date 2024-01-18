package main

import "fmt"

func main() {
	deferTest()
}

/*
defer 定义的函数会按照定义先后顺序倒序执行，类似于栈
*/
func deferTest() {

	defer fmt.Println("defer invoke")

	for i := 0; i < 5; i++ {
		defer fmt.Println("defer invoke", i)

	}
	return
}

/*
	执行结果：
defer invoke 4
defer invoke 3
defer invoke 2
defer invoke 1
defer invoke 0
defer invoke

*/
