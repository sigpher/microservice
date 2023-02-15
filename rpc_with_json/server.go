package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	serive := new(ServiceA)
	rpc.Register(serive)
	l, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	for {
		conn, _ := l.Accept()
		rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
