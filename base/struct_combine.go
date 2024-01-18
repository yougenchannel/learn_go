package main

type wheel struct {
	r int
	w int
}
type car struct {
	power int
	string
	int // 可以有多个不同类型的匿名属性

	wheel // 嵌入其他结构体

}

func main() {

}
