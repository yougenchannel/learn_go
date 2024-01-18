package main

import "fmt"

func main() {
	s := "Hello"

	sli := []byte(s)

	fmt.Println(len(sli))

	sli[0] = 's'

	fmt.Println(sli)

	newStr := string(sli)

	fmt.Println(newStr)
}

/*
运行结果：
5
[115 101 108 108 111]
sello

*/
