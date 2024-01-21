package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
)

func main() {
	var fileName = "io/test.txt"
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bytes))

	err = ioutil.WriteFile("io/wirte_file.txt", bytes, fs.ModeAppend)
	if err != nil {
		fmt.Println(err)
	}

}
