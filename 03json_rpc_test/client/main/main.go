package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main(){
	//conn, err := rpc.Dial("tcp", "localhost:8887")//原先go的协议是Gob
	conn, err := net.Dial("tcp", "localhost:8887")
	if err != nil{
		panic("连接失败")
	}
	var reply string
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	err = client.Call("HelloService.Hello", "flare", &reply)
	if err != nil{
		panic("调用失败")
	}
	fmt.Println(reply)
}
