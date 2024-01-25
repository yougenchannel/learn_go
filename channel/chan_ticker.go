package main

import (
	"fmt"
	"time"
)

func main() {

	rate := 10
	var dur time.Duration = time.Duration(1e9 / rate)
	ticker := time.Tick(dur)
	timer := time.NewTimer(1e19)

	for {
		select {
		case v := <-timer.C:
			fmt.Printf("timer", v)
		case v := <-ticker:
			fmt.Println("ticker", v)
		default:
			time.Sleep(1e9)
		}
	}

}
