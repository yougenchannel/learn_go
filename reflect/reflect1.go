package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := 5
	of := reflect.TypeOf(i)
	fmt.Println(of)

	fmt.Println(reflect.ValueOf(i))

	fmt.Println(reflect.Kind(i))

	i2 := reflect.ValueOf(i).Int()
	fmt.Println(i2)

	//fmt.Println(reflect.ValueOf(i).Bytes()) 运行时报错，panic: reflect: call of reflect.Value.Bytes on int Value
}
