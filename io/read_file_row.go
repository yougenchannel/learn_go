package main

import (
	"fmt"
	"io"
	"os"
)

const readFile = "io/test.txt"

func main() {
	open, err := os.Open(readFile)
	if err != nil {
		panic(err)
	}
	defer open.Close()

	var col1, col2 []string

	for {
		var v1, v2 string
		// 按照列进行读取
		_, err := fmt.Fscanln(open, &v1, &v2)
		if err == io.EOF {
			break
		}
		col1 = append(col1, v1)
		col2 = append(col2, v2)

	}

	fmt.Println(col1)
	fmt.Println(col2)

}
