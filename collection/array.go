package main

import (
	"fmt"
)

func main() {

	fun1()
	fun2()
	var a []int
	//var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println(len(a))
	var arrstr = [5]string{1: "Hel", 3: "shi"}
	var slice []string = arrstr[0:5]
	fmt.Println("slice len = ", len(slice))
	i := len(arrstr)
	fmt.Println(i)
	cap := cap(arrstr)
	fmt.Println(cap)
	for i, v := range arrstr {
		fmt.Println("idx = ", i, "value = ", v)
	}
}

func fun1() {
	a := [...]string{"st", "stt", "sstt"}

	for i, v := range a {
		fmt.Println("i = ", i, "value = ", v)
	}
}

func fun2() {
	var arr = new([5]int)
	arr[1] = 1
	fmt.Println(arr)
}
