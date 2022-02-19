package main

import (
	"net"
	"net/rpc"
)

type HelloService struct{

}

//远端处理逻辑函数
func (s *HelloService)Hello(request string, reply *string) error{
	//返回值通过修改reply的值
	*reply = "Hello, "+request
	return nil
}

func main(){
	//1.实例化一个server
	listener, _ := net.Listen("tcp", ":8887")
	//2.注册处理逻辑 handler
	_ = rpc.RegisterName("HelloService", &HelloService{})
	//3.启动服务
	conn, _ := listener.Accept()
	rpc.ServeConn(conn)//使用rpc对该请求服务

//	一连串代码除了rpc的作用域代码，其他都是net包的代码，可以再封装
//	几个问题：
//1.call id
//2.序列化和反序列化
//可以跨语言调用 1.go的rpc的序列化和反序列化的协议（Gob) 2.替换成常见的序列化json
}