package main

import (
	"flag"
	"fmt"
	"reflect"
)

var ngoroutine = flag.Int("n", 10000, "how many goroutines")

func main() {
	flag.Parse()
	leftmost := make(chan int)
	fmt.Println(reflect.TypeOf(leftmost))
	var left, right chan int = nil, leftmost
	for i := 0; i < *ngoroutine; i++ {
		left, right = right, make(chan int)
		go f(left, right)
	}
	right <- 0
	x := <-leftmost
	fmt.Println(x)
}

func f(left chan int, right chan int) {
	left <- 1 + <-right
}
