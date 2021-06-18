package main

import (
	"fmt"
	"net"
	"net/rpc"
	"time"
)

type Server struct{}

func (s *Server) Negate(i int64, reply *int64) error {
	*reply = -i
	return nil
}

func server() {
	rpc.Register(new(Server))
	ln, err := net.Listen("tcp", "localhost:19991")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		c, err := ln.Accept()
		if err != nil {
			continue
		}

		go rpc.ServeConn(c)
	}
}

func client() {
	c, err := rpc.Dial("tcp", "localhost:19991")
	if err != nil {
		fmt.Println(err)
		return
	}

	var result int64
	err = c.Call("Server.Negate", int64(999), &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Server.Negate(999) = ", result)
	}
}

func main() {
	go server()
	time.Sleep(1 * time.Second)
	go client()

	var input string
	fmt.Scanln(&input)
}
