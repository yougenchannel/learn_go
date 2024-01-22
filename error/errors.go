package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

var errNotFound error = errors.New("not found error")

func main() {
	err := getError()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func getError() error {

	reader := bufio.NewReader(os.Stdin)
	line, _, _ := reader.ReadLine()
	if len(line) < 100 {
		return fmt.Errorf("input string len less than 100, len = %d", len(line))
	}
	return nil
}
