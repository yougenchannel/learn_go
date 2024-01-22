package main

import (
	"fmt"
	"log"
)

func g() {
	panic("mock panic")
}

func test() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover panic")
		}

	}()
	g()
}

func main() {

	test()
	fmt.Println("programing continue run")
	for {

	}
}
