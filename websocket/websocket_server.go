package main

import (
	"fmt"
	websocket "golang.org/x/net/websocket"
	"log"
	"net/http"
)

func server(ws *websocket.Conn) {
	fmt.Println("new connection")
	buf := make([]byte, 100)
	buf = append(buf, []byte("sss")...)
	for {
		ws.Write(buf)
		//if _, err := ws.Read(buf); err != nil {
		//	fmt.Println("%s", err.Error())
		//	break
		//}

	}
	fmt.Println("close connection")
	ws.Close()
}

func main() {
	http.Handle("/websocket", websocket.Handler(server))
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("error websocket server")
	}

}
