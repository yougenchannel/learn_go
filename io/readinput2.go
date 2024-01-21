package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("please input something")
	// 以string的形式读取字符， 知道碰到delim设置的字符
	readString, err := reader.ReadString('\n')
	if err == nil {
		fmt.Println(readString)
		return
	}
	fmt.Println(err)
}
