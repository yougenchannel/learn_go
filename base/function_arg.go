package main

import "fmt"

func main() {

	function(1, callback)

	g := func() {
		sum := 0
		for i := 1; i <= 1e6; i++ {
			sum += i
		}
	}
	g()
	/*
		后面括号表示传参
	*/
	str := "aaa"
	go func(arg string) {

	}(str)

}

func function(x int, f func(int)) {
	f(x)
}

func callback(x int) {
	fmt.Println(x)

}

/*
匿名函数
*/
