package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	open, err := os.Open("io/testing.txt")
	if err != nil {
		return
	}
	defer open.Close()

	reader := bufio.NewReader(open)

	buffer := make([]byte, 1024)

	for {
		n, err := reader.Read(buffer)
		if n == 0 {
			return
		}
		if err == io.EOF {
			return
		}
		fmt.Println(" n = ", n, "bytes = ", string(buffer))
	}
}
