package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("io/test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()
	// File 是 io.Reader的实现
	reader := bufio.NewReader(file)
	for {
		// read string 以空格为分隔符每次读取
		// 也可以使用readLine来读取每行

		readString, err := reader.ReadString('\n')

		// isPrefix 用来返回是否为一行完全被读取， 如果缓冲区无法存下一行， 则后续的字符将在后面读取， 读取到最后一组后isPrefix返回false
		line, isPrefix, _ := reader.ReadLine()
		fmt.Println(isPrefix)
		fmt.Println(line)

		if err == io.EOF {
			fmt.Println(err)
			return
		}
		fmt.Printf("the input was : %s", readString)
	}

}
