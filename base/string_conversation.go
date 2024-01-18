package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	str := "12q"
	i, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(i)
}
