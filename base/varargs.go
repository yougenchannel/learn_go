package main

import "fmt"

func main() {
	myFunc("lisa", "bob", "paul")
}
func myFunc(args ...string) {
	fmt.Println(args)
	mySliceFunc(args)
}

/**
可变参数可通过slice 进行二次传递， 这里slice可以理解成List, 只是个可变数组
*/
func mySliceFunc(args []string) {
	fmt.Println(args)
}

/*
	运行结果
	[lisa bob paul]
	[lisa bob paul]
*/
