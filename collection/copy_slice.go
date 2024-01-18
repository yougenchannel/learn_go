package main

import "fmt"

func main() {
	sliFrom := []int{1, 2, 3}
	sliTo := make([]int, 10)
	n := copy(sliTo, sliFrom)
	fmt.Println("n = ", n)
	fmt.Println(sliTo)

	// 在slice末尾追加， 注意， 追加的索引从 len(slice) 开始向后追加， 如果cap不够会创建新的slice
	newSlice := append(sliTo, 4, 5, 6)
	fmt.Println(newSlice)

}

func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}
