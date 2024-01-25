package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"time"
)

func main() {
	ws, err := websocket.Dial("ws://localhost:8080/websocket", "", "http://localhost/")
	if err != nil {
		log.Fatal("client error")
	}
	go readFromServer(ws)
	time.Sleep(5e9)
}

func readFromServer(conn *websocket.Conn) {
	buf := make([]byte, 1000)
	for {

		if _, err := conn.Read(buf); err != nil {
			fmt.Printf("%s", err.Error())
			break
		}
		fmt.Println(string(buf), "----")
	}
}
