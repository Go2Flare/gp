package client_proxy

import (
	"go_code/gp/05new_helloworld/handler"
	"net/rpc"
)

type HelloServiceStub struct {
	*rpc.Client
}

//在go中没有类，对象，就意味着没有初始化方法
func NewHelloServiceClient(protol, address string) HelloServiceStub {
	conn, err := rpc.Dial(protol, address)
	if err != nil {
		panic("connect fail")
	}
	return HelloServiceStub{conn}
}

func (c *HelloServiceStub) Hello(request string, reply *string) error {
	err := c.Call(handler.HelloServiceName+".Hello", request, &reply)
	return err
}
