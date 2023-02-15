package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:8001")
	if err != nil {
		log.Fatal("rpc dial error: ", err)
	}
	args := &Args{10, 20}
	var reply int
	err = client.Call("ServiceA.Add", args, &reply)
	if err != nil {
		log.Fatal("serviceA.Add error: ", err)
	}
	fmt.Printf("ServiceA.Add: %d+%d=%d\n", args.X, args.Y, reply)

	var replay2 int
	divCall := client.Go("serviceA.Add", args, &replay2, nil)
	replyCall := <-divCall.Done
	fmt.Println(replyCall.Error)
	fmt.Println(replay2)
}
