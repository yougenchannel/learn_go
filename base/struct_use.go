package main

import (
	"fmt"
	"reflect"
)

type person struct {
	Name string `json:"name"`
	age  int    "a age tag"
}

func main() {
	p1 := new(person)
	p1.age = 1
	p1.Name = "yougen"

	// 指针类型， 其实和new出来的person一样
	p2 := &person{
		Name: "Martin",
		age:  0,
	}
	// 值类型
	p3 := person{
		Name: "Paul",
		age:  0,
	}

	t := reflect.TypeOf(p3)
	fmt.Println(t)
	fmt.Println(p2)

}
