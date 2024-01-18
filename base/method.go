package main

import (
	"fmt"
	"runtime"
	"strconv"
)

type TowInt struct {
	a, b int
}

// 非结构体
type IntVector []int

func main() {
	t := new(TowInt)
	t.a = 1
	t.b = 2
	fmt.Println(t.sum())
	fmt.Println(fmt.Sprint("%v", t))
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d", m.Alloc/1024)
}

func (this *TowInt) sum() int {
	return this.a + this.b
}

// 非结构体方法， 并且展示了不同接收器的同名方法
func (this IntVector) sum() (s int) {
	for _, v := range this {
		s += v
	}
	return s
}

func (this *TowInt) String() string {
	return "a = " + strconv.Itoa(this.a) + "b = " + strconv.Itoa(this.b)
}
