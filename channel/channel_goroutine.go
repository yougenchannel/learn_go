package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go sendData(ch)

	go getData(ch)

	// 一般只有输入端这么干， 输出端永远也不用关闭
	defer close(ch)
	ch1 := make(chan int)
	ch2 := make(chan int)
	go pump1(ch1)
	go pump1(ch2)
	go suck(ch1, ch2)
	time.Sleep(1e10)

}

func getData(ch chan string) {
	for {
		s, ok := <-ch
		// ok 判断channel是否被关闭
		// * 使用for range来读取，可以自动检测通道是否被关闭
		if !ok {
			return
		}
		fmt.Println("get data", s)
	}
}

func sendData(ch chan string) {
	ch <- "first"
	ch <- "second"
}

// 只读通道
func readChan(ch <-chan string) {
	for s := range ch {
		fmt.Println(s)
	}
}

// 只写通道
func writeChan(ch chan<- string) {
	for i := 0; i < 10; i++ {
		ch <- "chan"
	}
}

func pump1(ch chan int) {

	for i := 0; ; i += 2 {
		ch <- i
	}
}
func pump2(ch chan int) {
	for i := 0; ; i += 5 {
		ch <- i
	}

}

func suck(ch1, ch2 chan int) {
	for {
		select {
		case v := <-ch1:
			fmt.Println("received num2", v)
		case v := <-ch2:
			fmt.Println("received num5", v)
		default:
			fmt.Println("default: nothing ")
		}

	}
}
