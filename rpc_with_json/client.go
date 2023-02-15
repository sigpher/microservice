package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("dial error:", err)
	}
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	args := &Args{10, 20}
	var reply int
	client.Call("ServiceA.Add", args, &reply)
	if err != nil {
		log.Fatal("ServiceA.Add error:", err)
	}
	fmt.Printf("ServiceA.Add: %d+%d=%d\n", args.X, args.Y, reply)

	var reply2 int

	divCall := client.Go("ServiceA.Add", args, &reply2, nil)
	replyCall := <-divCall.Done
	fmt.Println(replyCall.Error)
	fmt.Println(reply2)
}
