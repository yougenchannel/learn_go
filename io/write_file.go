package main

import (
	"bufio"
	"fmt"
	"os"
)

const fileName = "io/write.txt"

func main() {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	writer := bufio.NewWriter(file)
	writeString := "Hello World!"

	_, err = writer.WriteString(writeString)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	// 记得刷新到文件
	err = writer.Flush()
	if err != nil {
		return
	}

}
