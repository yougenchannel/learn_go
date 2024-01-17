package main

import (
	"fmt"
)

const BEIJING = "BEIJING"

const (
	CHENGDU  = "CHENGDU"
	SHANGHAI = "SHANGHAI"
)

/*
	魔关键字iota
		iota 作为Go中的新特性， 通过行索引的方式（每换一行+1，可以弥补Go中没有没有枚举类的问题
*/

const (
	a = iota // 0
	b        // 1
	c        // 2
	d        // 3
)

const (
	e = 1 << iota
	f
	g
)

func main() {
	/*
	 只读， 修改编译器会报错
	*/

	const co int = 5

	fmt.Println(BEIJING)
	fmt.Println(co)

}
