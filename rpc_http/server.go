package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	service := new(ServiceA)
	rpc.Register(service)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Fatal("listen error :", err)
	}
	http.Serve(l, nil)
}
