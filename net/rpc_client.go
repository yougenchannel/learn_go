package main

import (
	"fmt"
	"learn_go/rpc_objects"
	"log"
	"net/rpc"
)

func main() {

	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("rpc client error")
	}
	var reply int
	args := &rpc_objects.Args{1, 2}
	err = client.Call("Args.Multiply", args, &reply)
	if err != nil {
		log.Fatal("rpc call fatal")
	}
	fmt.Println("success", reply)

}
