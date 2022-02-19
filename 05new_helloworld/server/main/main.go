package main

import (
	"go_code/gp/05new_helloworld/handler"
	serverProxy "go_code/gp/05new_helloworld/server_proxy"
	"net"
	"net/rpc"
)

func main(){
	listener, _ := net.Listen("tcp", ":8887")
	//_ = rpc.RegisterName(handler.HelloServiceName, &handler.HelloService{})
	serverProxy.ResgisterHelloService(&handler.NewHelloService{})
	for {
		conn, _ := listener.Accept()
		go rpc.ServeConn(conn)
	}
}