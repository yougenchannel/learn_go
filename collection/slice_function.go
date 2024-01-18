package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}                  //这是切片
	var arr = [3]string{"this", "is", "array"} // 这是数组
	funSlice2(arr[:])
	funSlice1(s)
	fmt.Println(s)
}

func funSlice1(slice []int) {
	slice[1] = 1
}
func funSlice2(slice []string) {
	slice[1] = "11"
}
