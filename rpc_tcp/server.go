package main

import (
	"log"
	"net"
	"net/rpc"
)

func main() {
	service := new(ServiceA)
	rpc.Register(service)

	l, err := net.Listen("tcp", "8001")
	if err != nil {
		log.Fatal("server error :", err)
	}
	for {
		conn, _ := l.Accept()
		rpc.ServeConn(conn)
	}
}
