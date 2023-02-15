package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:8001")
	if err != nil {
		log.Fatal("dailing: ", err)
	}

	// 同步调用
	args := &Args{10, 20}
	var reply int

	err = client.Call("ServiceA.Add", args, &reply)
	if err != nil {
		log.Fatal("Service.Add error :", err)
	}
	fmt.Printf("ServiceA.Add: %d + %d = %d\n", args.X, args.Y, reply)
	// 异步调用
	var reply2 int
	divCall := client.Go("ServiceA.Add", args, &reply2, nil)
	replyCall := <-divCall.Done
	fmt.Println(replyCall.Error)
	fmt.Println(reply2)
}
