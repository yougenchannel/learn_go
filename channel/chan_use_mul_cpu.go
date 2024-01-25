package main

import (
	"fmt"
	"time"
)

const NCPU int = 12

func DoAll() {
	sem := make(chan int, NCPU)
	for i := 0; i < NCPU; i++ {
		go DoPart(sem, i)

	}
	for i := 0; i < NCPU; i++ {
		fmt.Println(<-sem)
	}
}

func DoPart(sem chan int, i int) {
	sem <- i
}

func main() {
	DoAll()
	time.Sleep(1e10)
}
