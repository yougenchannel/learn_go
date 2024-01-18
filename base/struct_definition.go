package main

import (
	"bytes"
	"fmt"
	"sync"
)

type buffer struct {
	lock   sync.Mutex
	buffer bytes.Buffer
}

func main() {
	bf := new(buffer)
	bf.buffer.WriteByte(byte('s'))
	s := bf.buffer.String()
	fmt.Println(s)
}
