package main

import (
	"fmt"
	"os"
	"strings"
)

/*
读取命令行参数
*/
func main() {
	who := "Alice "
	if len(os.Args) > 1 {
		who += strings.Join(os.Args[1:], " ")
	}
	fmt.Println("Good Morning", who)
}
