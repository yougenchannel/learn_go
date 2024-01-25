package main

import (
	"learn_go/rpc_objects"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

const Addr = "localhost"

func main() {
	calc := new(rpc_objects.Args)
	rpc.Register(calc)
	rpc.HandleHTTP()
	listen, err := net.Listen("tcp", Addr+":1234")
	if err != nil {
		log.Fatal("error of rpc starting")
	}

	go http.Serve(listen, nil)
	time.Sleep(1e13)

}
