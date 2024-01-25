package main

import "fmt"

func main() {
	var testChan chan string
	testChan = make(chan string, 10)
	testChan <- "first message"

	s := <-testChan
	fmt.Printf(s)
}
