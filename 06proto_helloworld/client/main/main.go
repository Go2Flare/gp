package main

import (
	"fmt"
	"net/rpc"
)


func main(){
	client, err := rpc.Dial("tcp", "localhost:8887")
	if err != nil{
		panic("连接失败")
	}
	//var reply *string = new(string)
	var reply string
	err = client.Call("HelloService.Hello", "flare", &reply)
	//serviceMethod的前缀好处是，指明了函数的作用域
	//client.Hello("flare",&reply)//不用自己封装方法
	if err != nil{
		panic("调用失败")
	}
	fmt.Println(reply)
}
