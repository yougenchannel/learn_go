package main

import (
	"fmt"
)

/*
	声明变量四种方式：
*/

// 声明全局变量， 注意 d := 30， 这种方式不能用于定义全局变量
var t int
var u = 20
var v int = 30

func main() {

	var a int
	fmt.Println(a)

	var b int = 10
	fmt.Println(b)

	var c = 20
	fmt.Println(c)

	d := 30.1
	fmt.Print(d)

	// 单行声明多个变量
	var e, f int = 10, 11

	var g, h = 10, 3.14
	fmt.Println(e, f)
	fmt.Println(g, h)
	/* 多行变量声明
	Go语言中的基础数据类型包括以下几种：

	整型：int、int8、int16、int32、int64、uint、uint8、uint16、uint32、uint64等。其中int和uint类型的长度根据编译器位数而变化。
	浮点型：float32、float64。
	复数型：complex64、complex128。
	布尔型：bool，取值为true或false。
	字符串型：string。
	字节型：byte、rune。其中byte等同于uint8，rune等同于int32，用于表示UTF-8字符串的Unicode码点。
	*/
	var (
		i bool
		j float32

		k = 1.11
		l float64
		m int8
		n int16
		o string
		p int64
		q rune
		r byte
		s uint8
	)
	fmt.Println(i, j, k, l, m, n, o, p, q, r, s, t)

}
