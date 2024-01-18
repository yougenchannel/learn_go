package main

import (
	"fmt"
	"reflect"
	"sort"
)

type NotknownType struct {
	s1, s2, s3 string
}

func (n NotknownType) String() string {
	return n.s1 + " - " + n.s2 + " - " + n.s3
}

// variable to investigate:
var secret interface{} = NotknownType{"Ada", "Go", "Oberon"}

func main() {
	of := reflect.TypeOf(secret)

	fmt.Println(of)

	v := reflect.ValueOf(secret)
	fmt.Println(v)

	fmt.Println(v.Kind())

	for i := 0; i < v.NumField(); i++ {
		fmt.Println(i, v.Field(i))

	}

	result := v.Method(0).Call(nil)
	fmt.Println(result)

}

func minimumRemoval(beans []int) int64 {
	sort.Ints(beans)
	return 0
}
