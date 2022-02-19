package server_proxy

import (
	"go_code/gp/05new_helloworld/handler"
	"net/rpc"
)

type HelloServicer interface {
	// 实现了handler中NewHelloService的hello函数
	Hello(request string, reply *string) error
}

// ResgisterHelloService 如何做到和handler中对象的解耦？server关心的是函数->鸭子类型
//func ResgisterHelloService(srv handler.NewHelloService) {
//	rpc.RegisterName(handler.HelloServiceName, srv)
//}
//使用interface代替了这个对象
func ResgisterHelloService(srv HelloServicer) {
	rpc.RegisterName(handler.HelloServiceName, srv)
}
